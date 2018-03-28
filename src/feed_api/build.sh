#!/bin/bash


cluster_name=`echo "$1" | sed -r 's/^cop\.([^_\.]+)?_owt\.([^_\.]+)?_pdl\.([^_\.]+)?_cluster\.([^_\.]+)?.*/\4/'`
servicegroup_name=`echo "$1" | sed -r 's/^cop\.([^_\.]+)?_owt\.([^_\.]+)?_pdl\.([^_\.]+)?(.*)?_servicegroup\.([^_\.]+)?.*/\5/'`
service_name=`echo "$1" | sed -r 's/^cop\.([^_\.]+)?_owt\.([^_\.]+)?_pdl\.([^_\.]+)?(.*)?_service\.([^_\.]+)?.*/\5/'`
job_name=`echo "$1" | sed -r 's/^cop\.([^_\.]+)?_owt\.([^_\.]+)?_pdl\.([^_\.]+)?(.*)?_job\.([^_\.]+)?.*/\5/'`

#binary name
target=$2
default_target=feed_api

#cluster name
cluster=${cluster_name##*.}

#project path
project_path=$(cd `dirname $0`; pwd)
#src path
src_path=${project_path}
#release path
release_path=release
#bin path
release_bin_path=${release_path}/bin/
#config path
release_config_path=${release_path}/conf/

release_logs_path=${release_path}/logs/


if [ -d "src" ]; then
    printf "find src directory，user src directory \n"
    src_path=${project_path}/src
fi
if [ ! $target ]; then
    target=${default_target}
    printf "target is null,use default target name,%s \n" $target
fi
printEnv(){
    printf "Print Env \n"
    printf "============================================\n"
    printf "Commond Params        | %s %s \n" $1  $2
    printf "Project Path          | %s\n" $project_path
    printf "Src Path              | %s\n" $src_path
    printf "Target                | %s\n" $target
    printf "Service Nmae          | %s\n" $service_name
    printf "Cluster Name          | %s\n" $cluster_name
    printf "Job Name              | %s\n" $job_name
    printf "Service Group Name    | %s\n" $servicegroup_name

    printf "Release Path          | %s\n" $release_path
    printf "Release Bin  Path     | %s\n" $release_bin_path
    printf "Release Config Path   | %s\n" $release_config_path
    printf "============================================\n\n\n"
}

cleanDir(){
    printf "Clean Release Dir \n"
    printf "============================================\n"
    cd $project_path
    rm -rf $release_path
    if [ $? != 0 ]; then
        printf "Clean release dir failed\n"
        exit 101
    else
        printf "Clean release dir successed\n"
    fi

    mkdir -p $release_config_path
    mkdir -p $release_bin_path
    mkdir -p $release_logs_path
    printf "============================================\n\n\n"
}
buildBin(){
    printf "Build Bin \n"
    printf "============================================\n"
    cd $src_path
    printf "Pull dependence  ...\n"
    inkedep build
    if [ $? != 0 ]; then
        printf "Compiling project failed\n"
        exit 100
    fi
    printf "Pull dependence End\n"
    printf "Compiling project ...\n"
    go build -o $project_path/release/bin/$target
    if [ $? != 0 ]; then
        printf "Compiling project failed\n"
        exit 102
    else
        printf "Compiling project successed\n"
    fi
    cd $project_path
    printf "============================================\n\n\n"
}
copyConf(){
    printf "Copy Conf Files\n"
    printf "============================================\n"
    cd $project_path
    if [ -d "conf/$cluster" ]; then
        cp -r conf/$cluster/* $release_config_path
        echo "Copying conf/$cluster into release/conf"
    else
        cp conf/* $release_config_path
        echo "Copying conf into release/conf"
    fi

    if [ $? != 0 ]; then
        printf "Copying conf failed\n"
        exit 103
    fi
    printf "============================================\n\n\n"
}

copyInitScript() {
    printf "Copy Init Script\n"
    printf "============================================\n"
    cd $project_path

    cp start.sh $release_path
    echo "Copying init script into release/start.sh"

    if [ $? != 0 ]; then
        printf "Copying init script failed\n"
        exit 103
    fi
    printf "============================================\n\n\n"

}
printRelease(){
    printf "Print Release Directory\n"
    printf "============================================\n"
    cd $project_path
    find $release_path
    printf "============================================\n\n\n"
}

printEnv
cleanDir
buildBin
copyConf
copyInitScript
printRelease
exit 0
