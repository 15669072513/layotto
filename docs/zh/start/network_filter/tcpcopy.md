## Dump TCP����

### ����

��������quick start�ĵ�[ͨ��LayOtto����apollo��������](docs/zh/start/configuration/start-apollo.md)����demoʱ�����ܻ�ע�⵽�����ļ�config_apollo.json������ôһ�����ã�

```json
                {
                  "type": "tcpcopy",
                  "config": {
                    "port": "34904",
                    "strategy": {
                      "switch": "ON",
                      "interval": 30,
                      "duration": 10,
                      "cpu_max_rate": 80,
                      "mem_max_rate": 80
                    }
                  }
```
������õĺ���������ʱ����tcpcopy���������tcp����dump��

dump�����Ķ������������ݻ����� ${user's home directory}/logs/mosn Ŀ¼����/home/admin/logs/mosn Ŀ¼��:

![img.png](../../../../img/tcp_dump.png)

�����Խ���������ߺͻ�����ʩʹ����Щ���ݣ�������������طš���·��֤�ȡ�

### ������˵��

���ĵ�json�У�strategy��������Ҫ�������в����������ã���������˵�����£�

```go
type DumpConfig struct {
	Switch     string  `json:"switch"`       // dump switch.'ON' or 'OFF'
	Interval   int     `json:"interval"`     // dump sampling interval, unit: second
	Duration   int     `json:"duration"`     // Single sampling duration,unit: second
	CpuMaxRate float64 `json:"cpu_max_rate"` // cpu max rate.When cpu rate bigger than this threshold,dump function will be fused
	MemMaxRate float64 `json:"mem_max_rate"` // mem max rate.When memory rate bigger than this threshold,dump function will be fused
}
```

### ʵ��ԭ��

LayOtto������������MOSN�ϣ�ʹ��MOSN��filter��չ������������ĵ�tcpcopy��ʵ��MOSN��һ��network filter�����

�����Բο� [MOSN Դ����� - filter��չ����](https://mosn.io/blog/code/mosn-filters/) ʵ�����Լ���4��filter���