cd "`dirname $0`"
git add . -A
git commit -am "$1"
git push
/c/Python/python.exe ../go_appengine/appcfg.py --oauth2 update --no_cookies .
