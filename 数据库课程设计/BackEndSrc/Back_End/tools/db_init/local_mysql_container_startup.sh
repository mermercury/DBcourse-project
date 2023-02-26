ROOT_PASS=$1
DB_NAME=$2

if [ "${ROOT_PASS}" = "" ]; then
    ROOT_PASS="123456"
fi

if [ "${DB_NAME}" = "" ]; then
    DB_NAME="CourseDB"
fi
echo "ROOT_PASS=${ROOT_PASS},DBNAME=${DB_NAME}"
docker run -d -e MYSQL_ROOT_PASSWORD="${ROOT_PASS}" -e MYSQL_DATABASE="${DB_NAME}" -p 3306:3306 mysql