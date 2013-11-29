cd "`dirname $0`"
git add .
git commit -am "$1"
git push
/c/Python27/python.exe ../go_appengine/appcfg.py update --no_cookies .