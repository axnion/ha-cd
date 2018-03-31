# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
    config.vm.provider 'virtualbox' do |vb|
    end

    # Ansible Management Machine
    config.vm.define :mgmt do |mgmt|
        mgmt.vm.box = "bento/ubuntu-16.04"
        mgmt.vm.network :private_network, ip: "10.0.10.10"
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

    # Site A | Load Balancer
    config.vm.define :lb_a do |lb_a|
        lb_a.vm.box = "bento/ubuntu-16.04"
        lb_a.vm.network :private_network, ip: "10.0.10.20"
        lb_a.vm.network "forwarded_port", guest: 80, host: 8080
        lb_a.vm.network "forwarded_port", guest: 9000, host: 9000
        lb_a.vm.provider "virtualbox" do |vb|
            vb.memory = "512"
        end
    end

    # Site A | Application Server
    config.vm.define :app_a do |app_a|
        app_a.vm.box = "bento/ubuntu-16.04"
        app_a.vm.network :private_network, ip: "10.0.10.21"
        app_a.vm.provider "virtualbox" do |vb|
            vb.memory = "512"
        end
    end

    # Site A | Database
    config.vm.define :db_a do |db_a|
        db_a.vm.box = "bento/ubuntu-16.04"
        db_a.vm.network :private_network, ip: "10.0.10.22"
        db_a.vm.provider "virtualbox" do |vb|
            vb.memory = "1024"
        end
    end

#    # Site B | Load Balancer
#    config.vm.define :lb_b do |lb_b|
#        lb_b.vm.box = "bento/ubuntu-16.04"
#        lb_b.vm.network :private_network, ip: "10.0.10.30"
#        lb_b.vm.network "forwarded_port", guest: 80, host: 8083
#        lb_b.vm.provider "virtualbox" do |vb|
#            vb.memory = "1024"
#        end
#    end

    # Site B | Application Server
    config.vm.define :app_b do |app_b|
        app_b.vm.box = "bento/ubuntu-16.04"
        app_b.vm.network :private_network, ip: "10.0.10.31"
        app_b.vm.provider "virtualbox" do |vb|
            vb.memory = "1024"
        end
    end

    # Site B | Database
#    config.vm.define :db_b do |db_b|
#        db_b.vm.box = "bento/ubuntu-16.04"
#        db_b.vm.network :private_network, ip: "10.0.10.32"
#        db_b.vm.provider "virtualbox" do |vb|
#            vb.memory = "1024"
#        end
#    end
#
#
#    # Site C | Database
#    config.vm.define :db_c do |db_c|
#        db_c.vm.box = "bento/ubuntu-16.04"
#        db_c.vm.network :private_network, ip: "10.0.10.42"
#        db_c.vm.provider "virtualbox" do |vb|
#            vb.memory = "1024"
#        end
#    end
end
