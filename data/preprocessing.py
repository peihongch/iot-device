import os

import pandas as pd

# filename可以直接从盘符开始，标明每一级的文件夹直到csv文件，header=None表示头部为空，sep=' '表示数据间使用空格作为分隔符，如果分隔符是逗号，只需换成 ‘，’即可。
df = pd.read_csv('iot_telemetry_data.csv', sep=',')

# col_names = ["co", "humidity", "light", "lpg", "motion", "smoke", "temp"]
col_names = {"gas": ["co", "smoke", "lpg"], "th": ["temp", "humidity"], "light": ["light"], "motion": ["motion"]}
if not os.path.exists('devices'):
    os.mkdir('devices')
for k, v in col_names.items():
    # 根据device字段将数据拆分成多个设备的数据
    groups = df.groupby("device")
    for group in groups:
        if not os.path.exists('devices/%s' % group[0]):
            os.mkdir('devices/%s' % group[0])
        cols = ["ts", "device"]
        cols.extend(v)
        result = group[1][cols]
        result.to_csv('devices/%s/iot_telemetry_%s.csv' % (group[0], k), index=False,
                      header=True)
        print("wrote %d records to devices/%s/iot_telemetry_%s.csv" % (result.size, group[0], k))
