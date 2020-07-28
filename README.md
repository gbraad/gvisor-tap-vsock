# gvisor-tap-vsock 

## Build

```
make
```

## Run

### Host

#### Windows prerequisites

```
$service = New-Item -Path "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Virtualization\GuestCommunicationServices" -Name "00000400-FACB-11E6-BD58-64006A7986D3"
$service.SetValue("ElementName", "gvisor-tap-vsock")
```

More docs: https://docs.microsoft.com/en-us/virtualization/hyper-v-on-windows/user-guide/make-integration-service

#### Linux prerequisites

On Fedora 32, it worked out of the box. On others distros, you might have to look at https://github.com/mdlayher/vsock#requirements.

For CRC, the driver should be compiled with this patch: https://github.com/code-ready/machine-driver-libvirt/pull/45.

#### Run

```
(host) $ sudo bin/host -debug -logtostderr
```

### VM

```
(host) $ scp bin/vm crc:
(host) $ scp setup.sh crc:
(vm) $ sudo ./vm -debug -logtostderr [-windows if using windows]
(vm) $ sudo ./setup.sh
+ sudo ip addr add 192.168.127.0/24 dev O_O
+ sudo ip link set dev O_O up
+ sudo route del default gw 192.168.130.1
+ sudo route add default gw 192.168.127.1 dev O_O
(vm) $ ping -c1 192.168.127.1
(vm) $ curl http://redhat.com
```
