# k8s-homekit-cam

Docker image that runs a Go application as a sidecar container that tails the Home Assistant log file for HomeKit camera ffmpeg events, extracts the UDP port from the message, and dynamically creates a Kubernetes service for the camera RTP stream