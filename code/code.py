import subprocess
import json
import tempfile

def create_temp_file(event):
    temp_file = tempfile.NamedTemporaryFile(suffix=".go")
    with temp_file as fp:
        fp.write(bytes(event["body"], 'utf-8'))
    return temp_file

def run_go_code(temp_file):
    args = ["./bin/go", "version"]
    # args = ["go", "run", temp_file.name]
    popen = subprocess.Popen(args, stdout=subprocess.PIPE)
    popen.wait()


    args = ["./bin/go", "run", temp_file.name]
    # args = ["go", "run", temp_file.name]
    popen = subprocess.Popen(args, stdout=subprocess.PIPE)
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