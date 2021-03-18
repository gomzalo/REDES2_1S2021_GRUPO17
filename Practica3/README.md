# Practica 3
## Integrantes
Carnet | Nombre  
-- | -- 
201318652 | Gonzalo Antonio García Solares
201503584 | Eduardo Isaí Ajsivinac Xico
200413657 | Juan Pablo Renato Montúfar Chávez
200517803 | Johnny Mike Bravo Zamora
## Vodafone
Para Vodafone se utilizó una arquitectura hub and spoke, con un router principal y tres routers para distribución.
### Subneteo
Se crearon 8 subredes debido a que no se especifica el número de departamentos. Se tienen 2 host por cada departamento.
No Subred | Dirección de subred | Primera dirección asignable | Última dirección asignable | Dirección de broadcast | Máscara de subred
-- | -- | -- | -- | -- | --
1 | 192.168.77.0 | 192.168.77.1 | 192.168.77.30 | 192.168.77.31 | 255.255.255.224
2 | 192.168.77.32 | 192.168.77.33 | 192.168.77.62 | 192.168.77.63 | 255.255.255.224
3 | 192.168.77.64 | 192.168.77.65 | 192.168.77.94 | 192.168.77.95 | 255.255.255.224
4 | 192.168.77.96 | 192.168.77.97 | 192.168.77.126 | 192.168.77.127 | 255.255.255.224
5 | 192.168.77.128 | 192.168.77.129 | 192.168.77.158 | 192.168.77.159 | 255.255.255.224
6 | 192.168.77.160 | 192.168.77.161 | 192.168.77.190 | 192.168.77.191 | 255.255.255.224
7 | 192.168.77.192 | 192.168.77.193 | 192.168.77.222 | 192.168.77.223 | 255.255.255.224
8 | 192.168.77.224 | 192.168.77.225 | 192.168.77.254 | 192.168.77.255 | 255.255.255.224

### Tabla de IP's
Tipo  | Dirección IP | Máscara de Red | Gateway
 -- | -- | -- | --
Interfaz | 192.168.77.1 | 255.255.255.224 | 
Interfaz | 192.168.77.2 | 255.255.255.224 | 
Interfaz | 192.168.77.33 | 255.255.255.224 |
Interfaz | 192.168.77.34 | 255.255.255.224 |
Interfaz | 192.168.77.65 | 255.255.255.224 |
Interfaz | 192.168.77.66 | 255.255.255.224 |
Interfaz | 192.168.77.97 | 255.255.255.224 |
Host | 192.168.77.98 | 255.255.255.224 | 192.168.77.97
Host | 192.168.77.99 | 255.255.255.224 | 192.168.77.97
Host | 192.168.77.100 | 255.255.255.224 | 192.168.77.97
Interfaz | 192.168.77.129 | 255.255.255.224 |
Host | 192.168.77.130 | 255.255.255.224 | 192.168.77.129
Host | 192.168.77.131 | 255.255.255.224 | 192.168.77.129
Host | 192.168.77.132 | 255.255.255.224 | 192.168.77.129
Interfaz | 192.168.77.161 | 255.255.255.224 |
Host | 192.168.77.162 | 255.255.255.224 | 192.168.77.161
Host | 192.168.77.163 | 255.255.255.224 | 192.168.77.161
### Configuración

#### Configuración de las subredes
<pre><code>! Configuración del router PRINCIPAL
configure terminal
interface fastEthernet 0/1
no shutdown
ip address 192.168.77.1 255.255.255.224
exit
interface fastEthernet 1/0
no shutdown
ip address 192.168.77.33 255.255.255.224
exit
interface fastEthernet 2/0
no shutdown
ip address 192.168.77.65 255.255.255.224
end

! Configuración del router R2
configure terminal
interface fastEthernet 0/0
no shutdown
ip address 192.168.77.2 255.255.255.224
exit
interface fastEthernet 0/1
no shutdown
ip address 192.168.77.97 255.255.255.224
end

! Configuración del router R3
configure terminal
interface fastEthernet 0/0
no shutdown
ip address 192.168.77.34 255.255.255.224
exit
interface fastEthernet 0/1
no shutdown
ip address 192.168.77.129 255.255.255.224
end

! Configuración del router R3
configure terminal
interface fastEthernet 0/0
no shutdown
ip address 192.168.77.66 255.255.255.224
exit
interface fastEthernet 0/1
no shutdown
ip address 192.168.77.161 255.255.255.224
end</code>
</pre>

#### Configuración de ruteo
Se configuraron 6 subredes, dos por cada router.
<pre><code>! Ruteo en router PRINCIPAL
configure terminal
router ospf 555
network 192.168.77.0 0.0.0.31 area 555
network 192.168.77.32 0.0.0.31 area 555
network 192.168.77.64 0.0.0.31 area 555
exit

! Ruteo en router R2
configure terminal
router ospf 555
network 192.168.77.0 0.0.0.31 area 555
network 192.168.77.96 0.0.0.31 area 555
exit

! Ruteo en router R3
configure terminal
router ospf 555
network 192.168.77.32 0.0.0.31 area 555
network 192.168.77.128 0.0.0.31 area 555
exit

! Ruteo en router R4
configure terminal
router ospf 555
network 192.168.77.64 0.0.0.31 area 555
network 192.168.77.0 0.0.0.31 area 555
exit</code></pre>

#### Configuración de los host
<pre><code>! Red 1
! PC1
ip address 192.168.77.98 255.255.255.224 192.168.77.97

! PC2
ip address 192.168.77.99 255.255.255.224 192.168.77.97

! PC3
ip address 192.168.77.100 255.255.255.224 192.168.77.97


! Red 2
! PC4
ip address 192.168.77.130 255.255.255.224 192.168.77.129

! PC5
ip address 192.168.77.131 255.255.255.224 192.168.77.129

! PC6
ip address 192.168.77.132 255.255.255.224 192.168.77.129


! Red 3
! PC7
ip address 192.168.77.162 255.255.255.224 192.168.77.161

! PC8
ip address 192.168.77.163 255.255.255.224 192.168.77.161</code></pre>

## Telefónica
Para la configuración de la red de Telefónica, se utilizó una arquitectura de tres capas, con dos switch de acceso, dos routers de acceso, dos routers de distribución y un router núcleo.

### Subneteo
Se crearon 16 subredes, las cuales se dividen en 8 subredes con máscara 30 y 8 con máscara 27
No Subred | Dirección de subred | Primera dirección asignable | Última dirección asignable | Dirección de broadcast | Máscara de subred | Host/Subred
-- | -- | -- | -- | -- | -- | --
1 | 192.168.57.0 | 192.168.57.1 | 192.168.57.2 | 192.168.57.3 | 255.255.255.252 | 4
2 | 192.168.57.4 | 192.168.57.5 | 192.168.57.6 | 192.168.57.7 | 255.255.255.252 | 4
3 | 192.168.57.8 | 192.168.57.9 | 192.168.57.10 | 192.168.57.11 | 255.255.255.252 | 4 
4 | 192.168.57.12 | 192.168.57.13 | 192.168.57.14 | 192.168.57.15 | 255.255.255.252 | 4
5 | 192.168.57.16 | 192.168.57.17 | 192.168.57.18 | 192.168.57.19 | 255.255.255.252 | 4 
6 | 192.168.57.20 | 192.168.57.21 | 192.168.57.22 | 192.168.57.23 | 255.255.255.252 | 4 
7 | 192.168.57.24 | 192.168.57.25 | 192.168.57.26 | 192.168.57.27 | 255.255.255.252 | 4 
8 | 192.168.57.28 | 192.168.57.29 | 192.168.57.30 | 192.168.57.31 | 255.255.255.252 | 4 
9 | 192.168.57.32 | 192.168.57.33 | 192.168.57.62 | 192.168.57.63 | 255.255.255.224 | 32
10 | 192.168.57.64 | 192.168.57.65 | 192.168.57.94 | 192.168.57.95 | 255.255.255.224 | 32
11 | 192.168.57.96 | 192.168.57.97 | 192.168.57.126 | 192.168.57.127 | 255.255.255.224 | 32
12 | 192.168.57.128 | 192.168.57.129 | 192.168.57.158 | 192.168.57.159 | 255.255.255.224 | 32
13 | 192.168.57.160 | 192.168.57.161 | 192.168.57.190 | 192.168.57.191 | 255.255.255.224 | 32
14 | 192.168.57.192 | 192.168.57.193 | 192.168.57.222 | 192.168.57.223 | 255.255.255.224 | 32
15 | 192.168.57.224 | 192.168.57.225 | 192.168.57.254 | 192.168.57.225 | 255.255.255.224 | 32

### Tabla de IP's
Tipo  | Dirección IP | Máscara de Red | Gateway
 -- | -- | -- | -- 
Interfaz | 192.168.57.1 | 255.255.255.252 | 
Interfaz | 192.168.57.2 | 255.255.255.252 | 
Interfaz | 192.168.57.5 | 255.255.255.252 | 
Interfaz | 192.168.57.6 | 255.255.255.252 | 
Interfaz | 192.168.57.9 | 255.255.255.252 | 
Interfaz | 192.168.57.10 | 255.255.255.252 | 
Interfaz | 192.168.57.13 | 255.255.255.252 | 
Interfaz | 192.168.57.14 | 255.255.255.252 | 
Interfaz | 192.168.57.17 | 255.255.255.252 | 
Interfaz | 192.168.57.18 | 255.255.255.252 | 
Interfaz | 192.168.57.21 | 255.255.255.252 | 
Interfaz | 192.168.57.22 | 255.255.255.252 | 
Interfaz | 192.168.57.25 | 255.255.255.252 | 
Interfaz | 192.168.57.26 | 255.255.255.252 | 
Interfaz | 192.168.57.29 | 255.255.255.252 | 
Interfaz | 192.168.57.30 * | 255.255.255.252 | 
Interfaz | 192.168.57.33 | 255.255.255.224 | 
Host | 192.168.57.34 | 255.255.255.224 | 192.168.57.33 
Host | 192.168.57.35 | 255.255.255.224 | 192.168.57.33
Host | 192.168.57.36 | 255.255.255.224 | 192.168.57.33
Interfaz | 192.168.57.65 | 255.255.255.224 | 
Host | 192.168.57.66 | 255.255.255.224 | 192.168.57.65
Host | 192.168.57.67 | 255.255.255.224 | 192.168.57.65
Host | 192.168.57.68 | 255.255.255.224 | 192.168.57.65

\* Esta interfaz se utilizó para la conexión con la otra topología.

### Configuración de las subredes

<pre><code>! Configuración de NUCLEO
configure terminal
interface fastEthernet 0/1
ip address 192.168.57.1 255.255.255.252
no shutdown
exit
interface fastEthernet 1/0
ip address 192.168.57.5 255.255.255.252
no shutdown
end

! Configuración de DST1
configure terminal
interface fastEthernet 0/0
ip address 192.168.57.2 255.255.255.252
no shutdown
exit
interface fastEthernet 0/1
ip address 192.168.57.9 255.255.255.252
no shutdown
interface fastEthernet 1/0
ip address 192.168.57.13 255.255.255.252
no shutdown
exit
interface fastEthernet 2/0
ip address 192.168.57.25 255.255.255.252
no shutdown
end

! Configuración de DST2
configure terminal
interface fastEthernet 0/0
ip address 192.168.57.6 255.255.255.252
no shutdown
exit
interface fastEthernet 0/1
ip address 192.168.57.10 255.255.255.252
no shutdown
interface fastEthernet 1/0
ip address 192.168.57.17 255.255.255.252
no shutdown
exit
interface fastEthernet 2/0
ip address 192.168.57.21 255.255.255.252
no shutdown
end

! Configuración de ASC1
configure terminal
interface fastEthernet 0/0
ip address 192.168.57.14 255.255.255.252
no shutdown
exit
interface fastEthernet 0/1
ip address 192.168.57.22 255.255.255.252
no shutdown
interface fastEthernet 1/0
ip address 192.168.57.33 255.255.255.224
no shutdown
end

! Configuración de ASC2
configure terminal
interface fastEthernet 0/0
ip address 192.168.57.18 255.255.255.252
no shutdown
exit
interface fastEthernet 0/1
ip address 192.168.57.26 255.255.255.252
no shutdown
interface fastEthernet 1/0
ip address 192.168.57.65 255.255.255.224
no shutdown
end</code></pre>

### Configuración de ruteo.
<pre><code>! Configuración de NUCLEO
configure terminal
router rip
network 192.168.57.0
network 192.168.57.4
network 192.168.57.28
end

!Configuración de DST1
configure terminal
router rip
network 192.168.57.0
network 192.168.57.8
network 192.168.57.12
network 192.168.57.24
end

!Configuración de DST2
configure terminal
router rip
network 192.168.57.4
network 192.168.57.8
network 192.168.57.16
network 192.168.57.20
end

!Configuración de ASC1
configure terminal
router rip
network 192.168.57.12
network 192.168.57.20
network 192.168.57.32
end

!Configuración de ASC2
configure terminal
router rip
network 192.168.57.16
network 192.168.57.14
network 192.168.57.64
end</code></pre>

## Unión de las dos topologías por medio del protocolo EIGRP
Se configuró el protocolo de enrutamiento EIGRP para replicar las rutas de una topología a otra y así poder comunicarse entre sí, también se realizó la redisitribución de rutas de EIGRP - OSPF y EIGRP - RIP

### Configuración de las interfaces
<pre><code>! Configuración de PRINCIPAL
configure terminal
interface fastEthernet 0/0
ip address 192.168.57.29 255.255.255.252
no shutdown
end

! Configuración de NUCLEO
configure terminal
interface fastEthernet 0/0
ip address 192.168.57.30 255.255.255.252
no shutdown
end
</code></pre>

### Configuración del protocolo EIGRP
Este protocolo se configuró en los dos routers principales, colocando únicamente las rutas conectadas.

<pre><code>! Configuración de PRINCIPAL
configure terminal
router eigrp 12
no auto-summary
network 192.168.57.28 0.0.0.3
end

! Configuración de NUCLEO
configure terminal
router eigrp 12
no auto-summary
network 192.168.57.28 0.0.0.3
end</code></pre>

### Redistribución del protocolo OSPF en EIGRP
<pre><code>! Configuración de PRINCIPAL
configure terminal
router eigrp 12
! Significado de 1's
! 1. Metrica de ancho de banda en kbps
! 2. Delay, cada número significa 10 microsegundos
! 3. Fiabilidad, donde 255 es el 100% de confianza
! 4. Ancho de banda efectivo, donde 255 es el 100% cargado
! 5. Unidad máxima de transferencia de la ruta
redistribute ospf 555 metric 1 1 1 1 1
end</code></pre>

### Redistribución del protocolo RIP en EIGRP
<pre><code>! Configuración de NUCLEO
configure terminal
router eigrp 12
! Significado de ultimos numeros
! 1. -Numero de saltos a redistribuir
! 2. Delay, cada número significa 10 microsegundos
! 3. Fiabilidad, donde 255 es el 100% de confianza
! 4. Ancho de banda efectivo, donde 255 es el 100% cargado
! 5. Unidad máxima de transferencia de la ruta
redistribute rip metric 10000 100 255 1 1500
end</code></pre>

### Redistribución del protocolo EIGRP en OSPF
<pre><code>! Configuración de PRINCIPAL
configure terminal
router ospf 555
! Las subredes internas y externas
redistribute eigrp 12 metric external 1 external 2 internal
end</code></pre>

### Redistribución del protocolo EIGRP en RIP
<pre><code>! Configuración de NUCLEO
configure terminal
router rip
redistribute eigrp 12 metric 10
end</code></pre>

### Topología Final
![Topologia](https://github.com/gomzalo/REDES2_1S2021_P3_GRUPO17/blob/master/Images/TopologiaFinal.png "Topologia")
