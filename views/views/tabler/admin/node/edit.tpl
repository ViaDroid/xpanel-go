{{template "views/tabler/admin/header.tpl" .}}

<script src="//{{index .config "jsdelivr_url"}}/npm/jsoneditor@latest/dist/jsoneditor.min.js"></script>
<link href="//{{index .config "jsdelivr_url"}}/npm/jsoneditor@latest/dist/jsoneditor.min.css" rel="stylesheet" type="text/css">

<div class="page-wrapper">
    <div class="container-xl">
        <div class="page-header d-print-none text-white">
            <div class="row align-items-center">
                <div class="col">
                    <h2 class="page-title">
                        <span class="home-title">节点 #{{.node.Id}}</span>
                    </h2>
                    <div class="page-pretitle my-3">
                        <span class="home-subtitle">编辑节点信息</span>
                    </div>
                </div>
                <div class="col-auto ms-auto d-print-none">
                    <div class="btn-list">
                        <a id="save-node" href="#" class="btn btn-primary">
                            <i class="icon ti ti-device-floppy"></i>
                            保存
                        </a>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="page-body">
        <div class="container-xl">
            <div class="row row-deck row-cards">
                <div class="col-md-6 col-sm-12">
                    <div class="card">
                        <div class="card-header card-header-light">
                            <h3 class="card-title">基础信息</h3>
                        </div>
                        <div class="card-body">
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">名称</label>
                                <div class="col">
                                    <input id="name" type="text" class="form-control" value="{{.node.Name}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">连接地址</label>
                                <div class="col">
                                    <input id="server" type="text" class="form-control" value="{{.node.Server}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">IPv4地址</label>
                                <div class="col">
                                    <input type="text" class="form-control" value="{{.node.Ipv4}}" disabled>
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">IPv6地址</label>
                                <div class="col">
                                    <input type="text" class="form-control" value="{{.node.Ipv6}}" disabled>
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">流量倍率</label>
                                <div class="col">
                                    <input id="traffic_rate" type="text" class="form-control"
                                           value="{{.node.TrafficRate}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">接入类型</label>
                                <div class="col">
                                    <select id="sort" class="col form-select" value="{{.node.Sort}}">
                                        <option value="14" {{if czeq 14 .node.Sort}}selected{{end}}>Trojan</option>
                                        <option value="11" {{if eq 11 .node.Sort}}selected{{end}}>Vmess</option>
                                        <option value="2" {{if eq 2 .node.Sort}}selected{{end}}>TUIC</option>
                                        <option value="1" {{if eq 1 .node.Sort}}selected{{end}}>Shadowsocks2022</option>
                                        <option value="0" {{if eq 0 .node.Sort}}selected{{end}}>Shadowsocks</option>
                                    </select>
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">自定义配置</label>
                                <div id="custom_config"></div>
                                <label class="form-label col-form-label">
                                    请参考
                                    <a href="//wiki.sspanel.org/#/custom-config" target="_blank">
                                        wiki.sspanel.org/#/custom-config
                                    </a>
                                    修改节点自定义配置
                                </label>
                            </div>
                            <div class="form-group mb-3 row">
                                <span class="col">显示此节点</span>
                                <span class="col-auto">
                                    <label class="form-check form-check-single form-switch">
                                        <input id="type" class="form-check-input" type="checkbox" {{if .node.Type}}checked="" {{end}}>
                                    </label>
                                </span>
                            </div>
                            <div class="hr-text">
                                <span>动态倍率</span>
                            </div>
                            <div class="form-group mb-3 row">
                                <span class="col">启用动态流量倍率</span>
                                <span class="col-auto">
                                    <label class="form-check form-check-single form-switch">
                                        <input id="is_dynamic_rate" class="form-check-input" type="checkbox" {{if .node.IsDynamicRate}}checked="" {{end}}>
                                    </label>
                                </span>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">动态流量倍率计算方式</label>
                                <div class="col">
                                    <select id="dynamic_rate_type" class="col form-select" value="{{.node.DynamicRateType}}">
                                        <option value="0" {{if eq 0 .node.DynamicRateType}}selected{{end}}>Logistic</option>
                                        <option value="1" {{if eq 1 .node.DynamicRateType}}selected{{end}}>Linear</option>
                                    </select>
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">最大倍率</label>
                                <div class="col">
                                    <input id="max_rate" type="text" class="form-control" value="{{.node.MaxRate}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">最大倍率时间（时）</label>
                                <div class="col">
                                    <input id="max_rate_time" type="text" class="form-control" value="{{.node.MaxRateTime}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">最小倍率</label>
                                <div class="col">
                                    <input id="min_rate" type="text" class="form-control" value="{{.node.MinRate}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">最小倍率时间（时）</label>
                                <div class="col">
                                    <input id="min_rate_time" type="text" class="form-control" value="{{.node.MinRateTime}}">
                                </div>
                                <label class="form-label col-form-label">
                                    最大倍率时间必须大于最小倍率时间，否则将不会生效
                                </label>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-md-6 col-sm-12">
                    <div class="card">
                        <div class="card-header card-header-light">
                            <h3 class="card-title">其他信息</h3>
                        </div>
                        <div class="card-body">
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">等级</label>
                                <div class="col">
                                    <input id="node_class" type="text" class="form-control" value="{{.node.NodeClass}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">组别</label>
                                <div class="col">
                                    <input id="node_group" type="text" class="form-control" value="{{.node.NodeGroup}}">
                                </div>
                            </div>
                            <div class="hr-text">
                                <span>流量设置</span>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">已用流量</label>
                                <div class="col">
                                    <input id="node_bandwidth" type="text" class="form-control"
                                           value="{{.node.NodeBandwidth}}" disabled="">
                                </div>
                                <div class="col-auto">
                                    <button id="reset-bandwidth" class="btn btn-red">重置</button>
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">可用流量 (GB)</label>
                                <div class="col">
                                    <input id="node_bandwidth_limit" type="text" class="form-control"
                                           value="{{.node.NodeBandwidthLimit}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">流量重置日</label>
                                <div class="col">
                                    <input id="bandwidthlimit_resetday" type="text" class="form-control"
                                           value="{{.node.BandwidthlimitResetday}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">速率限制 (Mbps)</label>
                                <div class="col">
                                    <input id="node_speedlimit" type="text" class="form-control"
                                           value="{{.node.NodeSpeedlimit}}">
                                </div>
                            </div>
                            <div class="hr-text">
                                <span>高级选项</span>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">节点通讯密钥</label>
                                <input type="text" class="form-control" id="password" value="{{.node.Password}}"
                                       disabled="">
                                <div class="row my-3">
                                    <div class="col">
                                        <button id="reset-password" class="btn btn-red">重置</button>
                                        <button id="copy-password" class="copy btn btn-primary"
                                                data-clipboard-text="{{.node.Password}}">
                                            复制
                                        </button>
                                    </div>
                                </div>
                                <label class="form-label col-form-label">
                                    通讯密钥用于 WebAPI 节点模式鉴权，如需更改请点击重置
                                </label>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

<script>
    let clipboard = new ClipboardJS('.copy');
    clipboard.on('success', function (e) {
        $('#success-noreload-message').text('已复制到剪切板');
        $('#success-noreload-dialog').modal('show');
    });

    const container = document.getElementById('custom_config');
    let options = {
        modes: ['code', 'tree'],
    };
    const editor = new JSONEditor(container, options);
    editor.set(JSON.parse({{.node.CustomConfig}}))

    $("#reset-bandwidth").click(function () {
        $.ajax({
            url: '/admin/node/{{.node.Id}}/reset_bandwidth',
            type: 'POST',
            dataType: "json",
            success: function (data) {
                if (data.ret == 1) {
                    $('#success-message').text(data.msg);
                    $('#success-dialog').modal('show');
                } else {
                    $('#fail-message').text(data.msg);
                    $('#fail-dialog').modal('show');
                }
            }
        })
    });

    $("#reset-password").click(function () {
        $.ajax({
            url: '/admin/node/{{.node.Id}}/reset_password',
            type: 'POST',
            dataType: "json",
            success: function (data) {
                if (data.ret == 1) {
                    $('#success-message').text(data.msg);
                    $('#success-dialog').modal('show');
                } else {
                    $('#fail-message').text(data.msg);
                    $('#fail-dialog').modal('show');
                }
            }
        })
    });

    $("#save-node").click(function () {
        $.ajax({
            url: '/admin/node/{{.node.Id}}',
            type: 'PUT',
            dataType: "json",
            data: {
                {{range $key := .update_field}}
                {{$key}}: $('#{{$key}}').val(),
                {{end}}
                type: $("#type").is(":checked"),
                is_dynamic_rate: $("#is_dynamic_rate").is(":checked"),
                custom_config: JSON.stringify(editor.get()),
            },
            success: function (data) {
                if (data.ret == 1) {
                    $('#success-message').text(data.msg);
                    $('#success-dialog').modal('show');
                    window.setTimeout("location.href=top.document.referrer", {{index .config "jump_delay"}});
                } else {
                    $('#fail-message').text(data.msg);
                    $('#fail-dialog').modal('show');
                }
            }
        })
    });
</script>

{{template "views/tabler/admin/footer.tpl" .}}
