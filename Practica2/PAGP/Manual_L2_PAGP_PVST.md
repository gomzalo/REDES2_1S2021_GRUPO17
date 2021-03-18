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
conf t
hostname SERVIDOR
int e0/0
duplex full
no shutdown
exit
int range e1/0 - 3
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
int e0/0
switchport trunk encapsulation dot1q
switchport mode trunk
duplex full
exit
int range e1/0 - 3
switchport trunk encapsulation dot1q
switchport mode trunk
duplex full
end
wr

!*****   Po1   *****

conf t
int range e1/0 - 1
channel-protocol pagp
channel-group 1 mode desirable
end
wr

!*****   Po2   *****

conf t
int range e1/2 - 3
channel-protocol pagp
channel-group 2 mode desirable
end
wr

!*****  PVSTP    *****

conf t
spanning-tree mode pvst
end
wr

:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
::::::::::            ESW2 (CLIENTE1)            ::::::::::
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

ena

!*****   Interfaces   *****

!Ethernet Switch ESW2
conf t
hostname CLIENTE1
int range e1/0 - 3
duplex full
no shutdown
exit
int range e3/0 - 3
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
int range e1/0 - 3
switchport trunk encapsulation dot1q
switchport mode trunk
duplex full
end
wr

!*****   Access   *****

!VLAN-17

conf t
int range e3/0 - 1
switchport mode access
switchport access vlan 17
exit

!VLAN-27

int range e3/2 - 3
switchport mode access
switchport access vlan 27
end
wr

!*****   Po1   *****

conf t
int range e1/0 - 1
channel-protocol pagp
channel-group 1 mode desirable
end
wr

!*****   Po3   *****

conf t
int range e1/2 - 3
channel-protocol pagp
channel-group 3 mode desirable
end
wr

!*****  PVSTP    *****

conf t
spanning-tree mode pvst
end
wr

:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
::::::::::            ESW3 (CLIENTE0)            ::::::::::
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

ena

!*****   Interfaces   *****

!Ethernet Switch ESW3
conf t
hostname CLIENTE0
int range e1/0 - 3
duplex full
no shutdown
exit
int range e2/0 - 2
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
int range e1/0 - 3
switchport trunk encapsulation dot1q
switchport mode trunk
duplex full
exit
int range e2/0 - 2
switchport trunk encapsulation dot1q
switchport mode trunk
duplex full
end
wr

!*****   Po2   *****

conf t
int range e1/2 - 3
channel-protocol pagp
channel-group 2 mode desirable
end
wr

!*****   Po3   *****

conf t
int range e1/0 - 1
channel-protocol pagp
channel-group 3 mode desirable
end
wr

!*****   Po4   *****

conf t
int range e2/0 - 2
channel-protocol pagp
channel-group 4 mode desirable
end
wr

!*****  PVSTP    *****

conf t
spanning-tree mode pvst
end
wr

:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
::::::::::            ESW4 (CLIENTE2)            ::::::::::
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

ena

!*****   Interfaces   *****

!Ethernet Switch ESW4
conf t
hostname CLIENTE2
int range e0/0 - 1
duplex full
no shutdown
exit
int range e2/0 - 2
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
int range e0/0 - 1
switchport trunk encapsulation dot1q
switchport mode trunk
duplex full
exit
int range e2/0 - 2
switchport trunk encapsulation dot1q
switchport mode trunk
duplex full
end
wr

!*****   Po4   *****

conf t
int range e2/0 - 2
channel-protocol pagp
channel-group 4 mode desirable
end
wr

!*****  PVSTP    *****

conf t
spanning-tree mode pvst
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
set pcname ClsB2
ip 192.168.57.99 255.255.255.224 192.168.57.97
save

</code>
</pre>