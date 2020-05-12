import subprocess
import json

def lambda_handler(event, context):
    popen = subprocess.Popen("./go", stdout=subprocess.PIPE)
    popen.wait()


    return {
        'statusCode': 200,
        'body': json.dumps('Hello from Lambda!')
    }