import os

import pandas as pd

df = pd.read_csv('iot_telemetry_data.csv', sep=',')

col_names = {
    "gas": ["co", "smoke", "lpg"],
    "th": ["temp", "humidity"],
    "light": ["light"], "motion": ["motion"]
}
if not os.path.exists('devices'):
    os.mkdir('devices')
for k, v in col_names.items():
    # 根据device字段将数据拆分成多个设备的数据
    groups = df.groupby("device")
    idx = 1
    for group in groups:
        if not os.path.exists('devices/%d' % idx):
            os.mkdir('devices/%d' % idx)
        cols = ["ts", "device"]
        cols.extend(v)
        result = group[1][cols]
        result.to_csv('devices/%d/iot_telemetry_%s.csv' % (idx, k), index=False)
        print("wrote %d records to devices/%d/iot_telemetry_%s.csv" % (result.size, idx, k))
        idx += 1
