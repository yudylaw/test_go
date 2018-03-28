#/bin/bash
#go build
proj_name=`basename $(pwd)`
echo $proj_name
sudo mkdir -p /a8root/DSP-Server/$proj_name/logs
sudo mkdir -p /a8root/DSP-Server/$proj_name/bin
sudo mkdir -p /a8root/DSP-Server/$proj_name/conf
sudo cp conf/dev/* /a8root/DSP-Server/$proj_name/conf/
sudo cp $proj_name /a8root/DSP-Server/$proj_name/bin
sudo cp start.sh /a8root/DSP-Server/$proj_name
