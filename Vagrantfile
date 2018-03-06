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

    # Load Balancer
    config.vm.define :lb1 do |lb1|
        lb1.vm.box = "bento/ubuntu-16.04"
        lb1.vm.network :private_network, ip: "10.0.10.20"
        lb1.vm.network "forwarded_port", guest: 80, host: 8080
        lb1.vm.provider "virtualbox" do |vb|
            vb.memory = "512"
        end
    end

    # Web Server
    config.vm.define :web1 do |web1|
        web1.vm.box = "bento/ubuntu-16.04"
        web1.vm.network :private_network, ip: "10.0.10.30"
        web1.vm.network "forwarded_port", guest: 80, host: 8081
        web1.vm.provider "virtualbox" do |vb|
            vb.memory = "512"
        end
    end

    # Database
    config.vm.define :db1 do |db1|
        db1.vm.box = "bento/ubuntu-16.04"
        db1.vm.network :private_network, ip: "10.0.10.30"
        db1.vm.network "forwarded_port", guest: 80, host: 8082
        db1.vm.provider "virtualbox" do |vb|
            vb.memory = "512"
        end
    end
end
