SCRIPT_DIR=$(cd $(dirname $0); pwd)
cd $SCRIPT_DIR
curl -XDELETE localhost:9200/baito?pretty=true
curl -XPUT http://localhost:9200/baito?pretty -H 'Content-Type: application/json' -d @../json/index_settings.json 
curl -XPUT http://localhost:9200/baito/_mapping?pretty -H 'Content-Type: application/json' -d @../json/index_mappings.json