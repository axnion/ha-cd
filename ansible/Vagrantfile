# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
    config.vm.provider 'virtualbox' do |vb|
    end

    # Ansible Management Machine
    config.vm.define :mgmt do |mgmt|
        mgmt.vm.box = "bento/ubuntu-16.04"
        mgmt.vm.network :private_network, ip: "10.0.10.5"
        mgmt.vm.provider "virtualbox" do |vb|
            vb.memory = "512"
        end
        mgmt.vm.provision "shell", privileged: true, inline: <<-SHELL
          #! /bin/bash
          apt-get -y install software-properties-common
          apt-add-repository -y ppa:ansible/ansible
          apt-get update
          apt-get -y install ansible
          chown -R vagrant:vagrant /vagrant
          echo cd /vagrant >> .bashrc
        SHELL
    end

# LOAD BALANCERS ---------------------------------------------------------------

    # Site A | Load Balancer
    config.vm.define :lb_a do |lb_a|
        lb_a.vm.box = "bento/ubuntu-16.04"
        lb_a.vm.network :private_network, ip: "10.0.10.10"
        lb_a.vm.network "forwarded_port", guest: 80, host: 8081
        lb_a.vm.provider "virtualbox" do |vb|
            vb.memory = "512"
        end
    end

    # Site B | Load Balancer
#    config.vm.define :lb_b do |lb_b|
#        lb_b.vm.box = "bento/ubuntu-16.04"
#        lb_b.vm.network :private_network, ip: "10.0.10.20"
#        lb_b.vm.network "forwarded_port", guest: 80, host: 8082
#        lb_b.vm.provider "virtualbbx" do |vb|
#            vb.memory = "512"
#        end
#    end
#
#    # Site C | Load Balancer
#    config.vm.define :lb_c do |lb_c|
#        lb_c.vm.box = "bento/ubuntu-16.04"
#        lb_c.vm.network :private_network, ip: "10.0.10.30"
#        lb_c.vm.network "forwarded_port", guest: 80, host: 8083
#        lb_c.vm.provider "virtualbox" do |vb|
#            vb.memory = "512"
#        end
#    end
#
#    # Site D | Load Balancer
#    config.vm.define :lb_d do |lb_d|
#        lb_d.vm.box = "bento/ubuntu-16.04"
#        lb_d.vm.network :private_network, ip: "10.0.10.40"
#        lb_d.vm.network "forwarded_port", guest: 80, host: 8084
#        lb_d.vm.provider "virtualbbx" do |vb|
#            vb.memory = "512"
#        end
#    end

# APPLICATION SERVER -----------------------------------------------------------

    # Site A | Application Server
    config.vm.define :app_a do |app_a|
        app_a.vm.box = "bento/ubuntu-16.04"
        app_a.vm.network :private_network, ip: "10.0.10.11"
        app_a.vm.provider "virtualbox" do |vb|
            vb.memory = "512"
        end
    end

    # Site B | Application Server
    config.vm.define :app_b do |app_b|
        app_b.vm.box = "bento/ubuntu-16.04"
        app_b.vm.network :private_network, ip: "10.0.10.21"
        app_b.vm.provider "virtualbox" do |vb|
            vb.memory = "512"
        end
    end

    # Site C | Application Server
    config.vm.define :app_c do |app_c|
        app_c.vm.box = "bento/ubuntu-16.04"
        app_c.vm.network :private_network, ip: "10.0.10.31"
        app_c.vm.provider "virtualbox" do |vb|
            vb.memory = "512"
        end
    end

    # Site D | Application Server
    config.vm.define :app_d do |app_d|
        app_d.vm.box = "bento/ubuntu-16.04"
        app_d.vm.network :private_network, ip: "10.0.10.41"
        app_d.vm.provider "virtualbox" do |vb|
            vb.memory = "512"
        end
    end

# DATABASE SERVER --------------------------------------------------------------

    # Site A | Database
    config.vm.define :db_a do |db_a|
        db_a.vm.box = "bento/ubuntu-16.04"
        db_a.vm.network :private_network, ip: "10.0.10.12"
        db_a.vm.provider "virtualbox" do |vb|
            vb.memory = "1024"
        end
    end

    # Site B | Database
    config.vm.define :db_b do |db_b|
        db_b.vm.box = "bento/ubuntu-16.04"
        db_b.vm.network :private_network, ip: "10.0.10.22"
        db_b.vm.provider "virtualbox" do |vb|
            vb.memory = "1024"
        end
    end

    # Site C | Database
    config.vm.define :db_c do |db_c|
        db_c.vm.box = "bento/ubuntu-16.04"
        db_c.vm.network :private_network, ip: "10.0.10.32"
        db_c.vm.provider "virtualbox" do |vb|
            vb.memory = "1024"
        end
    end

    # Site D | Database
    config.vm.define :db_d do |db_d|
        db_d.vm.box = "bento/ubuntu-16.04"
        db_d.vm.network :private_network, ip: "10.0.10.42"
        db_d.vm.provider "virtualbox" do |vb|
            vb.memory = "1024"
        end
    end

end
