package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/nxadm/tail"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const log_start = "Start FFmpeg with"
const log_port = "rtcpport="
const port_digits = 5
const service_name_prefix = "homekit-cam"

func main() {
	fmt.Println("Starting k8s-homekit-cam...")

	flagCleanupPtr := flag.String("cleanup", "", "delete K8s services older than DURATION")
	flagMonitorPtr := flag.String("monitor", "", "log file to monitor")
	flag.Parse()

	flagCleanup := *flagCleanupPtr
	flagMonitor := *flagMonitorPtr

	if len(flagCleanup) > 0 {
		duration, err := time.ParseDuration(flagCleanup)
		if err != nil {
			panic(err.Error())
		}
		cleanup_services(duration)
	}

	if len(flagMonitor) > 0 {
		tail_file(flagMonitor)
	}
}

func tail_file(filename string) {
	fmt.Println("Tailing file", filename)

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace_bytes, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		panic(err.Error())
	}
	namespace := string(namespace_bytes)

	// Create a tail
	t, err := tail.TailFile(
		filename, tail.Config{Follow: true, ReOpen: true, Location: &tail.SeekInfo{Offset: 0, Whence: io.SeekEnd}})
	if err != nil {
		panic(err)
	}

	// Iterate each received line
	for line := range t.Lines {
		if strings.Contains(line.Text, log_start) &&
			strings.Contains(line.Text, log_port) {
			port_idx := strings.Index(line.Text, log_port)
			port := line.Text[port_idx+len(log_port) : port_idx+len(log_port)+port_digits]
			port_int, err := strconv.Atoi(port)
			if err != nil {
				panic(err)
			}
			fmt.Println("Creating K8s service for FFmpeg port:", port)

			clientset.CoreV1().Services(namespace).Create(context.TODO(), &apiv1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-%s", service_name_prefix, port),
					Namespace: namespace,
					Labels: map[string]string{
						"app": "homeassistant",
					},
				},
				Spec: apiv1.ServiceSpec{
					Type: apiv1.ServiceTypeNodePort,
					Selector: map[string]string{
						"app": "homeassistant",
					},
					Ports: []apiv1.ServicePort{
						apiv1.ServicePort{
							Protocol: apiv1.ProtocolUDP,
							Port:     int32(port_int),
							NodePort: int32(port_int),
						},
					},
					ExternalTrafficPolicy: apiv1.ServiceExternalTrafficPolicyTypeLocal,
				},
			},
				metav1.CreateOptions{})
		}
	}
}

func cleanup_services(max_service_age time.Duration) {
	fmt.Println("Deleting services older than", max_service_age)

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	namespace_bytes, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		panic(err.Error())
	}
	namespace := string(namespace_bytes)

	services, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, value := range services.Items {
		if  time.Since(value.CreationTimestamp.Time) > max_service_age &&
			strings.HasPrefix(value.Name, service_name_prefix) {
			fmt.Println("Deleting aged out camera service", value.Name, "created", time.Since(value.CreationTimestamp.Time).Minutes(), "minutes ago")
			err := clientset.CoreV1().Services(namespace).Delete(context.TODO(), value.Name, metav1.DeleteOptions{})
			if err != nil {
				fmt.Println("Error deleting service", value.Name, ":", err.Error())
			}		
		}
	}
}