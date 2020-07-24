#!/bin/bash

echo "GRANT ALL ON \`$MYSQL_DATABASE\`.* TO '$MYSQL_USER'@'%' IDENTIFIED BY '$MYSQL_PASSWORD';" | mysql -hlocalhost -uroot -p"$MYSQL_ROOT_PASSWORD"
