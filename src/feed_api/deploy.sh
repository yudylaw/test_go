#/bin/bash
go build
proj_name=`basename $(pwd)`
DEPLOY_GIT="git@code.inke.cn:video/deploy/feed_api.git"
TMP_GIT_DIR="git_dir"
DATETIME=`date "+%H_%M_%m_%d_%Y"`
MSG_GIT=${proj_name}

# 开始部署
echo "deploy to git starting...."
mkdir -p ${TMP_GIT_DIR}/bin
mkdir -p ${TMP_GIT_DIR}/conf
#拉取代码
cd ${TMP_GIT_DIR};git checkout master;git pull;cd ..
mv ./$proj_name ${TMP_GIT_DIR}/bin
cp -r conf/deploy/* ${TMP_GIT_DIR}/conf
cp start.sh ${TMP_GIT_DIR}/bin
cd ${TMP_GIT_DIR}
git add .;git commit -am "${MSG_GIT} ${DATETIME} auto deploy to staging";git push -u origin master
echo "deploy to git ok"
