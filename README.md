<div aling="center">

# Observabilidade - Stack de Monitoramento e Análise


<img src="./assets/logo_gova-project.png" width="800" height="400" alt= "logo de um cachorro com oculos">
</div>

<div aling="center">
  

## :8ball: Descrição

Este repositório contém os arquivos e instruções necessários para configurar uma stack de observabilidade, utilizando Prometheus, Node Exporter, Grafana e uma aplicação em Go que simula um petshop. O objetivo dessa atividade é coletar métricas, criar dashboards e alertas para monitorar e analisar o desempenho da aplicação bem como implantar todo ecossistema criado localmente em um servidor remoto na AWS em produção utilizando o Ansible.


## :smile: Status do Projeto

![→_-em desenvolvimento-green](https://user-images.githubusercontent.com/87483916/224333868-804e4712-ecb7-492b-8c47-6c6bff8e3b7a.svg)


## 🚀 Começando

As instruções a seguir são destinadas a ambientes de teste e desenvolvimento. Para implantação em produção, consulte a seção "Implantação em Produção" para obter um guia passo a passo específico. 

<Br>

### 📋 As tecnologias de pré-requisitos:

<Br>

* Docker: É necessário ter o Docker instalado em seu ambiente para executar os componentes da stack em containers isolados. O Docker permite a criação, execução e gerenciamento de containers de aplicativos. 
<Br>

* Docker Compose: O Docker Compose é uma ferramenta que facilita a especificação e execução de vários containers em conjunto. Ele permite definir a configuração e as dependências dos containers em um arquivo YAML, simplificando o processo de execução da stack de observabilidade.
<Br>

* Linux: A atividade e as instruções fornecidas pressupõem o uso de um sistema operacional Linux. Certifique-se de ter um ambiente Linux configurado e funcional para executar as etapas descritas.
<Br>

* Ansible: O Ansible é uma ferramenta de automação que permite configurar e provisionar recursos de maneira consistente em servidores remotos. Será utilizado para criar um playbook Ansible que automatizará a implantação da stack de observabilidade em um servidor remoto AWS.
<Br>

* Acesso a um servidor remoto AWS: É necessário ter acesso a um servidor remoto na plataforma AWS (Amazon Web Services) para implantar a stack de observabilidade em produção. Isso envolve a configuração de um servidor remoto na AWS e o acesso a ele usando chaves SSH públicas.

<Br>

### 🔧 Instalação

#### :anger: Clonando o repositório:
<hr>

1. Clone este repositório em sua máquina local utilizando o seguinte comando:


```
git clone https://github.com/seu-usuario/nome-do-repositorio.git
```

####  Instalando o Docker
<hr>
1.Instale o Docker em sua máquina utilizando o script oficial Docker


```
$ sudo curl -fsSL https://get.docker.com -o get-docker.sh
yes | sudo sh get-docker.sh 

```
2. Instale o Docker Compose 


```
$ sudo apt-get install docker-compose

```

####  Executar o Docker Compose
<hr>
 
 1.Navegue até o diretório do repositório clonado em sua máquina e execute o Docker Compose com o seguinte comando:


```
docker-compose up -d

```

2.Esse comando iniciará os containers do Prometheus e Grafana em segundo plano. Os containers serão criados e configurados automaticamente com as imagens do Docker

####  Acessar o serviço do Prometheus

####  Configurar o Grafana



## :airplane: Roadmap
 

- [ ] **Preparação e Configuração Inicial**
  - [ ] Instalar o Docker em sua máquina local.
  - [ ] Configurar acesso à AWS e criar uma instância EC2 para o servidor remoto.
  - [ ] Configurar chaves SSH para acessar a instância remota.
  - [ ] Conectar-se ao servidor remoto e garantir que esteja pronto para receber a implantação.

- [ ] **Configuração da Stack de Observabilidade Localmente**
  - [ ] Clonar o repositório da stack de observabilidade em seu ambiente local.
  - [ ] Configurar o arquivo docker-compose.yml para definir os serviços Prometheus, Node Exporter, Grafana e a aplicação do petshop.
  - [ ] Executar o Prometheus via Docker e configurar o mapeamento de portas.
  - [ ] Executar o Node Exporter e configurar sua execução no docker-compose.yml.
  - [ ] Executar a aplicação do petshop via docker-compose e configurar o acesso às métricas no Prometheus.
  - [ ] Executar o Grafana via Docker e configurar os Data Sources para acessar as métricas do Prometheus.
  - [ ] Criar dashboards no Grafana para visualizar as métricas do petshop e do Node Exporter.

- [ ] **Implantação Automatizada na AWS**
  - [ ] Criar um playbook do Ansible para configurar a instância EC2 com as dependências necessárias.
  - [ ] Configurar o playbook do Ansible para instalar o Docker e o Docker Compose na instância remota.
  - [ ] Transferir o arquivo docker-compose.yml para a instância remota.
  - [ ] Executar o playbook do Ansible para implantar a stack de observabilidade na instância remota.

- [ ] **Teste e Monitoramento em Produção**
  - [ ] Acessar a instância remota e verificar se a stack de observabilidade foi implantada corretamente.
  - [ ] Monitorar as métricas coletadas pelo Prometheus e visualizá-las no Grafana.
  - [ ] Testar a geração de alertas com base nas métricas configuradas.
  - [ ] Realizar ajustes e otimizações com base nos dados de observabilidade coletados.

- [ ] **Documentação Oficial**
  - [ ] Criar uma documentação oficial para o repositório, incluindo instruções de instalação, configuração e uso da stack de observabilidade.
  - [ ] Incluir links úteis para recursos relacionados, como documentação do Prometheus, Grafana e Ansible.
  - [ ] Acrescentar exemplos de consultas Prometheus e configurações avançadas no Grafana.
  - [ ] Garantir que a documentação esteja atualizada e seja de fácil acesso para os usuários.

Lembre-se de marcar as tarefas concluídas com [x] à medida que você avança no roadmap.




## 📦 Implantação
  
  <Br>
  
  https://github.com/rc-ventura/opencms-ec2-tomcat-nginx-deployment/assets/87483916/be8c65b1-750e-4ce2-bf0e-45ce6536cab3

  
  #### :crystal_ball: Docker
    
  
  ![----_-Em Construção  ](https://github.com/rc-ventura/opencms-ec2-tomcat-nginx-deployment/assets/87483916/249d676a-b9dd-48bf-ba4e-a9811bb59b2f)

    
  #### :mag: Terraform 
  
 ![----_-Em Construção  ](https://github.com/rc-ventura/opencms-ec2-tomcat-nginx-deployment/assets/87483916/7025d250-223d-40e0-9cae-2284ec976905)

 #### :computer: OpenCms

  ![----_-Em Construção  ](https://github.com/rc-ventura/opencms-ec2-tomcat-nginx-deployment/assets/87483916/b24d7d98-f411-45a0-8fe1-a15357b10529)

 #### ::palm_tree::  Nginx e Tomcat

![----_-Em Construção  ](https://github.com/rc-ventura/opencms-ec2-tomcat-nginx-deployment/assets/87483916/6d14b2f8-1bcd-4234-8c39-c3a7d7626cc6)

 #### ::leaves::  Postgresql
    
 ![----_-Em Construção  ](https://github.com/rc-ventura/opencms-ec2-tomcat-nginx-deployment/assets/87483916/6d14b2f8-1bcd-4234-8c39-c3a7d7626cc6)



<Br>


## :book: Documentação
  
##### <a href="https://docs.aws.amazon.com/"> * Amazon Web Services (AWS)</a> 
##### <a href="https://https://docs.docker.com/"> * Docker</a>
##### <a href="https://docs.docker.com/compose/"> * Docker Compose </a>
##### <a href="https://www.linux.org/"> * Linux </a>
##### <a href=" https://prometheus.io/docs/"> * Prometheus</a>
##### <a href="https://grafana.com/docs/"> * Grafana </a>
##### <a href="https://docs.ansible.com/"> * Ansible </a>
##### <a href="https://github.com/prometheus/node_exporter"> * Node Exporter </a>

## :crystal_ball: Imagens Docker

#####  * harbor.sflabs.com.br/processo/petshop:latest (Petshop) - imagem com hospedagem privada
##### <a href="https://hub.docker.com/r/prom/prometheus"> * prom/prometheus</a>
##### <a href="https://hub.docker.com/r/grafana/grafana"> * grafana/grafana </a>
##### <a href="https://hub.docker.com/r/prom/node-exporter"> * prom/node-exporter </a>


## 🛠️ Tecnologias

 <div  style="display: inline-block" >

 <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/amazonwebservices/amazonwebservices-original-wordmark.svg" width="80"  /> 
 <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" width="80">
 <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/terraform/terraform-original.svg" width="80" />
 <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/ansible/ansible-original.svg" width="80" />
 <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/prometheus/prometheus-original.svg" width="80" />
 <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/grafana/grafana-original.svg" width="80"  />
 <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/linux/linux-original.svg" width="80" />
          
          
                

   
</div>


## 🖇️ Colaborando

Por favor, leia o [COLABORACAO.md](https://gist.github.com/usuario/linkParaInfoSobreContribuicoes) para obter detalhes sobre o nosso código de conduta e o processo para nos enviar pedidos de solicitação.

## 📌 Versão

 Para as versões disponíveis, observe as [tags neste repositório](https://github.com/suas/tags/do/projeto). 

## ✒️ Autores

Mencione todos aqueles que ajudaram a levantar o projeto desde o seu início

* **Rafael Ventura** - *Responsável técnico* - [Dev Ventura](https://github.com/rc-ventura)

Você também pode ver a lista de todos os [colaboradores](https://github.com/usuario/projeto/colaboradores) que participaram deste projeto.

## 📄 Licença

Este projeto está sob a licença (sua licença) - veja o arquivo [LICENSE.md](https://github.com/usuario/projeto/licenca) para detalhes.

## 🎁 Expressões de gratidão


---
⌨️ com ❤️ por [Rafael Ventura](https://gist.github.com/rc-ventura) 😊
