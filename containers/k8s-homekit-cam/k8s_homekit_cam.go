package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/nxadm/tail"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)


const log_start = "Start FFmpeg with"
const log_port = "rtcpport="
const port_digits = 5

func main() {
	fmt.Println("Starting hass_homekit_cam_k8s...")

	if len(os.Args) < 2 {
		panic("Log file name missing from command line arguments")
	}

	log_filename := os.Args[1]

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
		log_filename, tail.Config{Follow: true, ReOpen: true, Location: &tail.SeekInfo{Offset: 0, Whence: io.SeekEnd}})
	if err != nil {
		panic(err)
	}

	// Print the text of each received line
	for line := range t.Lines {
		//fmt.Println(line.Text)
		if strings.Contains(line.Text, log_start) &&
			strings.Contains(line.Text, log_port) {
			port_idx := strings.Index(line.Text, log_port)
			port := line.Text[port_idx+len(log_port) : port_idx+len(log_port)+port_digits]
			port_int, err := strconv.Atoi(port)
			if err != nil {
				panic(err)
			}
			fmt.Println("Got port:", port)

			clientset.CoreV1().Services(namespace).Create(context.TODO(), &apiv1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("homekit-cam-%s", port),
					Namespace: namespace,
					Labels: map[string]string{
						"app": "homeassistant",
					},
					Annotations: map[string]string{
						"metallb.universe.tf/allow-shared-ip": "homekit",
					},
				},
				Spec: apiv1.ServiceSpec{
					Type:           apiv1.ServiceTypeLoadBalancer,
					LoadBalancerIP: "192.168.0.37",
					Selector: map[string]string{
						"app": "homeassistant",
					},
					Ports: []apiv1.ServicePort{
						apiv1.ServicePort{
							Protocol: apiv1.ProtocolUDP,
							Port:     int32(port_int),
						},
					},
				},
			},
				metav1.CreateOptions{})
		}
	}
}
