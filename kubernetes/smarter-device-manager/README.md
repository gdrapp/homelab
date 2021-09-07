# smarter-device-manager

Enables k8s containers to access devices (linux device drivers) available on nodes

## Installation

```bash
kubectl apply -k kubernetes/smarter-device-manager
```

## Configuration

Add devices via regular expression to the configmap.yaml file. All matched devices are added as resources.

## Troubleshooting

The node will show the devices it recognizes as resources in the node object in Kubernetes. The example below shows a raspberry PI.

```bash
kubectl describe node pike5


Name:               pike5
Roles:              <none>
Labels:             beta.kubernetes.io/arch=arm
                    beta.kubernetes.io/os=linux
                    fake-image-generator-rpi=enabled
                    smarter-device-manager=enabled
Annotations:        node.alpha.kubernetes.io/ttl: 0
CreationTimestamp:  Mon, 02 Dec 2019 09:22:56 -0600
Taints:             <none>
Unschedulable:      false
Lease:
  HolderIdentity:  <unset>
  AcquireTime:     <unset>
  RenewTime:       <unset>
Conditions:
  Type             Status  LastHeartbeatTime                 LastTransitionTime                Reason                       Message
  ----             ------  -----------------                 ------------------                ------                       -------
  MemoryPressure   False   Thu, 16 Jan 2020 08:20:06 -0600   Mon, 02 Dec 2019 09:22:56 -0600   KubeletHasSufficientMemory   kubelet has sufficient memory available
  DiskPressure     False   Thu, 16 Jan 2020 08:20:06 -0600   Wed, 04 Dec 2019 09:47:08 -0600   KubeletHasNoDiskPressure     kubelet has no disk pressure
  PIDPressure      False   Thu, 16 Jan 2020 08:20:06 -0600   Mon, 02 Dec 2019 09:22:56 -0600   KubeletHasSufficientPID      kubelet has sufficient PID available
  Ready            True    Thu, 16 Jan 2020 08:20:06 -0600   Mon, 16 Dec 2019 14:58:05 -0600   KubeletReady                 kubelet is posting ready status. AppArmor enabled
Addresses:
  InternalIP:  XXX.XXX.XXX.XXX
  Hostname:    pike5
Capacity:
  cpu:                        4
  ephemeral-storage:          14999512Ki
  memory:                     873348Ki
  pods:                       110
  smarter-devices/gpiochip0:  10
  smarter-devices/gpiochip1:  10
  smarter-devices/gpiochip2:  10
  smarter-devices/gpiomem:    10
  smarter-devices/i2c-1:      10
  smarter-devices/snd:        10
  smarter-devices/vchiq:      10
  smarter-devices/vcs:        0
  smarter-devices/vcsm:       10
  smarter-devices/vcsm-cma:   0
  smarter-devices/video10:    0
  smarter-devices/video11:    0
  smarter-devices/video12:    0
  smarter-devices/video4:     10
Allocatable:
  cpu:                        4
  ephemeral-storage:          13823550237
  memory:                     770948Ki
  pods:                       110
  smarter-devices/gpiochip0:  10
  smarter-devices/gpiochip1:  10
  smarter-devices/gpiochip2:  10
  smarter-devices/gpiomem:    10
  smarter-devices/i2c-1:      10
  smarter-devices/snd:        10
  smarter-devices/vchiq:      10
  smarter-devices/vcs:        0
  smarter-devices/vcsm:       10
  smarter-devices/vcsm-cma:   0
  smarter-devices/video10:    0
  smarter-devices/video11:    0
  smarter-devices/video12:    0
  smarter-devices/video4:     10
System Info:
  Machine ID:                 XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
  System UUID:                XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
  Boot ID:                    XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
  Kernel Version:             5.3.0-1014-raspi2
  OS Image:                   Ubuntu 19.10
  Operating System:           linux
  Architecture:               arm
  Container Runtime Version:  docker://19.3.2
  Kubelet Version:            v1.13.5
  Kube-Proxy Version:         v1.13.5
Non-terminated Pods:          (5 in total)
  Namespace                   Name                              CPU Requests  CPU Limits  Memory Requests  Memory Limits  AGE
  ---------                   ----                              ------------  ----------  ---------------  -------------  ---
  argus                       smarter-device-manager-gdmjk      10m (0%)      100m (2%)   15Mi (1%)        15Mi (1%)      43d
Allocated resources:
  (Total limits may be over 100 percent, i.e., overcommitted.)
  Resource                   Requests     Limits
  --------                   --------     ------
  cpu                        560m (14%)   850m (21%)
  memory                     365Mi (48%)  365Mi (48%)
  ephemeral-storage          0 (0%)       0 (0%)
  smarter-devices/gpiochip0  0            0
  smarter-devices/gpiochip1  0            0
  smarter-devices/gpiochip2  0            0
  smarter-devices/gpiomem    0            0
  smarter-devices/i2c-1      0            0
  smarter-devices/snd        0            0
  smarter-devices/vchiq      1            1
  smarter-devices/vcs        0            0
  smarter-devices/vcsm       1            1
  smarter-devices/vcsm-cma   0            0
  smarter-devices/video10    0            0
  smarter-devices/video11    0            0
  smarter-devices/video12    0            0
  smarter-devices/video4     2            2
Events:                      <none>
```
