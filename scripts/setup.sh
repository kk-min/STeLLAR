#!/bin/bash
sudo apt-get update

# For the client to run unattended
sudo apt-get install --no-install-recommends --assume-yes tmux

# Install JDK 11 and Gradle for compiling ZIP archive for Java functions
sudo apt-get install --no-install-recommends --assume-yes openjdk-11-jdk
wget https://services.gradle.org/distributions/gradle-8.1.1-bin.zip -P /tmp
sudo unzip -d /opt/gradle /tmp/gradle-8.1.1-bin.zip
sudo ln -s /opt/gradle/gradle-8.1.1 /opt/gradle/latest
echo "export GRADLE_HOME=/opt/gradle/latest" >> /etc/profile.d/gradle.sh
echo "export PATH=\${GRADLE_HOME}/bin:\${PATH}" >> /etc/profile.d/gradle.sh
sudo chmod +x /etc/profile.d/gradle.sh
source /etc/profile.d/gradle.sh

# Equivalent to "ulimit -n 1000000", see https://superuser.com/questions/1289345/why-doesnt-ulimit-n-work-when-called-inside-a-script
sudo sh -c "echo \"* soft nofile 1000000\" >> /etc/security/limits.conf"
sudo sh -c "echo \"* hard nofile 1000000\" >> /etc/security/limits.conf"
sudo sh -c "echo \"root soft nofile 1000000\" >> /etc/security/limits.conf"
sudo sh -c "echo \"root hard nofile 1000000\" >> /etc/security/limits.conf"

sudo mkdir ../latency-samples
echo "Please run: ulimit -n 8192"
echo "Recommended: tmux new -s stellar"
echo "AWS example run: sudo ./main -o latency-samples -g endpoints -c experiments/tests/aws/data-transfer.json"
