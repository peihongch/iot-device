import os

import pandas as pd

# filename可以直接从盘符开始，标明每一级的文件夹直到csv文件，header=None表示头部为空，sep=' '表示数据间使用空格作为分隔符，如果分隔符是逗号，只需换成 ‘，’即可。
df = pd.read_csv('iot_telemetry_data.csv', sep=',')

col_names = ["co", "humidity", "light", "lpg", "motion", "smoke", "temp"]
if not os.path.exists('devices'):
    os.mkdir('devices')
for i in range(len(col_names)):
    col = col_names[i]

    # 根据device字段将数据拆分成多个设备的数据
    groups = df.groupby("device")
    for group in groups:
        if not os.path.exists('devices/%s' % group[0]):
            os.mkdir('devices/%s' % group[0])
        result = group[1][["ts", "device", col]]
        result.to_csv('devices/%s/iot_telemetry_%s.csv' % (group[0], col), index=False,
                      header=False)
        print("wrote %d records to devices/%s/iot_telemetry_%s.csv" % (result.size, group[0], col))
