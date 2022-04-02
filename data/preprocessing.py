import os

import pandas as pd

# filename可以直接从盘符开始，标明每一级的文件夹直到csv文件，header=None表示头部为空，sep=' '表示数据间使用空格作为分隔符，如果分隔符是逗号，只需换成 ‘，’即可。
df = pd.read_csv('iot_telemetry_data.csv', sep=',')

col_names = ["co", "humidity", "light", "lpg", "motion", "smoke", "temp"]
if not os.path.exists('devices'):
    os.mkdir('devices')
for i in range(len(col_names)):
    col = col_names[i]
    df[["ts", "device", col]].to_csv('devices/iot_telemetry_%s.csv' % col, index=False, header=False)
