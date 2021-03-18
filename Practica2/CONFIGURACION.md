## Comandos de configuración
El siguiente documento, contiene los comandos para configurar cada dispositivo.

### Ether-Switches L3

<pre>
<code>

:::::::::::::::::::::::::::::::::::::::::::::::::::::::::
::::::::::            ESW1 (SERVER)            ::::::::::
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::

ena

!*****   Interfaces   *****

!Ethernet Switch ESW1 
configure terminal
hostname SERVIDOR
interface range fastEthernet 1/0 - 4
speed 100
duplex full
no shutdown
end
wr

!*****   VTP   *****

conf t
vtp domain Grupo17
vtp password Grupo17
vtp mode server
end
wr

!*****   VLAN   *****

!::::::::::         VLAN 17 - Administracion         ::::::::::

conf t
vlan 17
name Administracion17
end
wr

!::::::::::         VLAN 27 - Profesores         ::::::::::

conf t
vlan 27
name Profesores27
end
wr

!::::::::::         VLAN 37 - Clase A         ::::::::::

conf t
vlan 37
name ClaseA37
end
wr

!::::::::::         VLAN 47 - Clase B         ::::::::::

conf t
vlan 47
name ClaseB47
end
wr

!*****   INTERVLAN   *****

!::::::::::         VLAN17        ::::::::::

conf t
int vlan 17
ip address 192.168.57.1 255.255.255.224
no shu
exit

!::::::::::         VLAN27        ::::::::::

int vlan 27
ip address 192.168.57.33 255.255.255.224
no shu
exit

!::::::::::         VLAN37        ::::::::::

conf t
int vlan 37
ip address 192.168.57.65 255.255.255.224
no shu
exit

!::::::::::         VLAN47        ::::::::::

int vlan 47
ip address 192.168.57.97 255.255.255.224
no shu
end

!::::::::::         Ruteo        ::::::::::

conf t
ip routing
end
wr

!*****   Trunk   *****

conf t
int range fa1/0 - 4
switchport mode trunk
speed 100
duplex full
end
wr

!*****   Po1   *****

conf t
int range f1/0 - 1
channel-group 1 mode on
end
wr

!*****   Po2   *****

conf t
int range f1/2 - 3
channel-group 2 mode on
end
wr

:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
::::::::::            ESW2 (CLIENTE1)            ::::::::::
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

ena

!*****   Interfaces   *****

!Ethernet Switch ESW2
configure terminal
hostname CLIENTE1
interface range fastEthernet 1/0 - 3
speed 100
duplex full
no shutdown
exit
interface range fastEthernet 1/12 - 15
speed 100
duplex full
no shutdown
end
wr

!*****   VTP   *****

conf t
vtp domain Grupo17
vtp password Grupo17
vtp mode client
end
wr

!*****   Trunk   *****

conf t
int range fa1/0 - 3
switchport mode trunk
speed 100
duplex full
end
wr

!*****   Access   *****

!VLAN-17

conf t
int fa1/12
switchport mode access
switchport access vlan 17
exit
int fa1/14
switchport mode access
switchport access vlan 17
exit

!VLAN-27

int fa1/13
switchport mode access
switchport access vlan 27
exit
int fa1/15
switchport mode access
switchport access vlan 27
end
wr

!*****   Po1   *****

conf t
int range f1/0 - 1
channel-group 1 mode on
end
wr

!*****   Po3   *****

conf t
int range f1/2 - 3
channel-group 3 mode on
end
wr

:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
::::::::::            ESW3 (CLIENTE0)            ::::::::::
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

ena

!*****   Interfaces   *****

!Ethernet Switch ESW3
configure terminal
hostname CLIENTE0
interface range fastEthernet 1/0 - 6
speed 100
duplex full
no shutdown
end
wr

!*****   VTP   *****

conf t
vtp domain Grupo17
vtp password Grupo17
vtp mode client
end
wr

!*****   Trunk   *****

conf t
int range fa1/0 - 6
switchport mode trunk
speed 100
duplex full
end
wr

!*****   Po2   *****

conf t
int range f1/0 - 1
channel-group 2 mode on
end
wr

!*****   Po3   *****

conf t
int range f1/2 - 3
channel-group 3 mode on
end
wr

!*****   Po4   *****

conf t
int range f1/4 - 6
channel-group 4 mode on
end
wr

:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
::::::::::            ESW4 (CLIENTE2)            ::::::::::
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

ena

!*****   Interfaces   *****

!Ethernet Switch ESW4
configure terminal
hostname CLIENTE2
interface range fastEthernet 1/0 - 4
speed 100
duplex full
no shutdown
end
wr

!*****   VTP   *****

conf t
vtp domain Grupo17
vtp password Grupo17
vtp mode client
end
wr

!*****   Trunk   *****

conf t
int range fa1/0 - 4
switchport mode trunk
duplex full
speed 100
end
wr

!*****   Po4   *****

conf t
int range f1/0 - 2
channel-group 4 mode on
end
wr

</code>
</pre>

### Switches L2

<pre>
<code>

:::::::::::::::::::::::::::::::::::::::::::::::::
::::::::::             SW1             ::::::::::
:::::::::::::::::::::::::::::::::::::::::::::::::

ena

!*****   Interfaces   *****

!Switch L2 SW1
configure terminal
hostname SW1
interface ethernet 0/0
duplex full
no shutdown
exit
interface range ethernet 1/0 - 1
duplex full
no shutdown
end
wr

!*****   Trunk   *****

configure terminal
interface ethernet 0/0
switchport trunk encapsulation dot1q
switchport mode trunk
duplex full
end
wr

!*****   Access   *****

conf t
int range e1/0 - 1
switchport mode access
switchport access vlan 47
end
wr

:::::::::::::::::::::::::::::::::::::::::::::::::
::::::::::             SW2             ::::::::::
:::::::::::::::::::::::::::::::::::::::::::::::::

ena

!*****   Interfaces   *****

!Switch L2 SW2
configure terminal
hostname SW2
interface ethernet 0/0
duplex full
no shutdown
exit
interface range ethernet 1/0 - 1
duplex full
no shutdown
end
wr

!*****   Trunk   *****

configure terminal
interface ethernet 0/0
switchport trunk encapsulation dot1q
switchport mode trunk
duplex full
end
wr

!*****   Access   *****

conf t
int range e1/0 - 1
switchport mode access
switchport access vlan 37
end
wr

:::::::::::::::::::::::::::::::::::::::::::::::::
::::::::::             SW3             ::::::::::
:::::::::::::::::::::::::::::::::::::::::::::::::

ena

!*****   Interfaces   *****

!Switch L2 SW3
configure terminal
hostname SW3
interface ethernet 0/0
duplex full
no shutdown
exit
interface range ethernet 1/0 - 1
duplex full
no shutdown
end
wr

!*****   Trunk   *****

configure terminal
interface ethernet 0/0
switchport trunk encapsulation dot1q
switchport mode trunk
duplex full
end
wr

!*****   Access   *****

!VLAN-17

conf t
int e1/0
switchport mode access
switchport access vlan 17
exit

!VLAN-27

int e1/1
switchport mode access
switchport access vlan 27
end
wr

</code>
</pre>

### VPCs

<pre>
<code>

:::::::::::::::::::::::::::::::::::::::::::::::::
::::::::::            IP               ::::::::::
:::::::::::::::::::::::::::::::::::::::::::::::::

*****   ServidorAdmin   *****

!ServidorAdmin
set pcname SrvAdm
ip 192.168.57.2 255.255.255.224 192.168.57.1
save

*****   Administracion1   *****

!Administracion1
set pcname Admin1
ip 192.168.57.3 255.255.255.224 192.168.57.1
save

*****   Administracion2   *****

!Administracion2
set pcname Admin2
ip 192.168.57.4 255.255.255.224 192.168.57.1
save

*****   ServidorProfesor   *****

!ServidorProfesor
set pcname SvProf
ip 192.168.57.34 255.255.255.224 192.168.57.33
save

*****   Profesor1   *****

!Profesor1
set pcname Prof1
ip 192.168.57.35 255.255.255.224 192.168.57.33
save

*****   Profesor2   *****

!Profesor2
set pcname Prof2
ip 192.168.57.36 255.255.255.224 192.168.57.33
save

*****   ClaseA1   *****

!ClaseA1
set pcname ClsA1
ip 192.168.57.66 255.255.255.224 192.168.57.65
save

*****   ClaseA2   *****

!ClaseA2
set pcname ClsA2
ip 192.168.57.67 255.255.255.224 192.168.57.65
save

*****   ClaseB1   *****

!ClaseB1
set pcname ClsB1
ip 192.168.57.98 255.255.255.224 192.168.57.97
save

*****   ClaseB2   *****

!ClaseB2
set pcname ClsA1
ip 192.168.57.99 255.255.255.224 192.168.57.97
save

</code>
</pre>