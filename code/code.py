import subprocess
import json
import tempfile

def create_temp_file(event):
    temp_file = tempfile.NamedTemporaryFile(suffix=".go")
    with temp_file as fp:
        fp.write(bytes(event["body"]))
    return temp_file

def run_go_code(temp_file):
    args = ["./../bin/go", "run", temp_file]
    popen = subprocess.Popen(args, stdout=subprocess.PIPE)
    popen.wait()


def lambda_handler(event, context):
    temp_file = create_temp_file(event)
    run_go_code(temp_file)

    return {
        'statusCode': 200,
        'body': json.dumps('Hello from Lambda!')
    }