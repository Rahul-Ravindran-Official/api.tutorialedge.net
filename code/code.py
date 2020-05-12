import subprocess
import json
import tempfile
import os
import tarfile

def create_temp_file(event):
    tempFile = tempfile.NamedTemporaryFile(delete=False, dir="/tmp", suffix=".go")
    with tempFile as fp:
        tempFile.write(bytes(event["body"], 'utf-8'))

    return tempFile

def run_go_code(temp_file):
    print(os.listdir())

    if not os.path.exists("/tmp/go"):
        go_code = tarfile.open("./code/go.tar.gz", "r:gz")
        go_code.extractall("/tmp")
        go_code.close()
    
    print(os.listdir("/tmp/go"))

    my_env = os.environ.copy()
    my_env["PATH"] = "/usr/sbin:/sbin:/tmp/go/bin"
    my_env["GOROOT"] = "/tmp/go"
    my_env["GOPATH"] = "/tmp"

    args = ["go", "version"]
    popen = subprocess.Popen(args, stdout=subprocess.PIPE, env=my_env)
    popen.wait()
        
    args = ["go", "run", temp_file.name]
    popen = subprocess.Popen(args, stdout=subprocess.PIPE, env=my_env)
    popen.wait()


def lambda_handler(event, context):
    temp_file = create_temp_file(event)
    run_go_code(temp_file)

    return {
        'statusCode': 200,
        'body': json.dumps('Hello from Lambda!')
    }

# def main():
#     event = {}
#     event["body"] = """package main

#     import "fmt"

#     func main() {

#         mymap := make(map[string]int)

#         mymap["elliot"] = 25

#         // we can use this if statement to check to see if 
#         // a given key "elliot" exists within a map in Go
#         if _, ok := mymap["elliot"]; ok {
#             // the key 'elliot' exists within the map
#             fmt.Println(mymap["elliot"])
#         }
#     }""" 

#     status = lambda_handler(event, None)
#     print(status)

# if __name__ == '__main__':
#     main()