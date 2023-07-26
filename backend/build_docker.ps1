# docker build -f formatter/Dockerfile -t mukmookk/backend/formatter:latest .
# docker push mukmookk/backend/formatter:latest

$LANGS = "python", "golang", "java", "javascript"
foreach ($LANG in $LANGS)
{
    docker build -f language_servers/"$LANG"/Dockerfile -t mukmookk/backend/lang-servers-"$LANG":latest .
    docker push mukmookk/backend/lang-servers-"$LANG":latest
}
