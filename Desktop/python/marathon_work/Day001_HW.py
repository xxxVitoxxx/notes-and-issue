from urllib.request import urlretrieve
import os
os.makedirs('./daaa',exist_ok=True) #exists_ok = true →如果該資料夾存在，不執行不會噴error

files = []
if 'Homework.txt' in files:
    print('[O] , Homework.txt is exists!!')
else:
    print('[X] , Homework.txt isn\'t exists!!')

with open('./daaa/Homework.txt','w') as file:
    file.write('Hello World')

with open('./daaa/Homework.txt','r') as file2:
    data = file2.read()

if len('Hello world') == len(data):
    print('[O] 檢查 Homework.txt 檔案字數是否符合 Hello World 字數')
else:
    print('[X] 檢查 Homework.txt 檔案字數是否符合 Hello World 字數')