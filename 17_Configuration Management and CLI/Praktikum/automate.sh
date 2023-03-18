#!/bin/bash

nama=$1
facebook=$2
linkedin=$3

today=$(date)
hostname=$(uname -n)
hostdevice=$(uname -srvmpio)

maindir="${nama} at ${today}"


mkdir -p "${maindir}"/{about_me/{personal,professional},my_friends,my_system_info}

echo "https://www.facebook.com/${facebook}" > facebook.txt
mv -f facebook.txt "${maindir}"/about_me/personal/facebook.txt

echo "https://www.linkedin.com/in/${linkedin}" > linkedin.txt
mv -f linkedin.txt "${maindir}"/about_me/professional/linkedin.txt

curl https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt > list_of_friends.txt
mv -f list_of_friends.txt "${maindir}"/my_friends/list_of_friends.txt

echo "My username: ${hostname}"> about_this_laptop.txt
echo "With host: ${hostdevice}" >> about_this_laptop.txt

mv -f about_this_laptop.txt "${maindir}"/my_system_info/about_this_laptop.txt

ping google.com -c 3 >> internet_connection.txt
mv -f internet_connection.txt "${maindir}"/my_system_info/internet_connection.txt

tree
