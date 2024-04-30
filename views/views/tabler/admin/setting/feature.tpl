{{template "views/tabler/admin/header.tpl" .}}

<div class="page-wrapper">
    <div class="container-xl">
        <div class="page-header d-print-none text-white">
            <div class="row align-items-center">
                <div class="col">
                    <h2 class="page-title">
                        <span class="home-title">其他设置</span>
                    </h2>
                    <div class="page-pretitle my-3">
                        <span class="home-subtitle">设置站点的其他设置</span>
                    </div>
                </div>
                <div class="col-auto ms-auto d-print-none">
                    <div class="btn-list">
                        <a id="save-setting" href="#" class="btn btn-primary">
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
                <div class="col-md-12">
                    <div class="card">
                        <div class="card-header">
                            <ul class="nav nav-tabs card-header-tabs" data-bs-toggle="tabs">
                                <li class="nav-item">
                                    <a href="#display" class="nav-link active" data-bs-toggle="tab">功能显示</a>
                                </li>
                                <li class="nav-item">
                                    <a href="#log" class="nav-link" data-bs-toggle="tab">用户日志</a>
                                </li>
                                <li class="nav-item">
                                    <a href="#checkin" class="nav-link" data-bs-toggle="tab">签到</a>
                                </li>
                            </ul>
                        </div>
                        <div class="card-body">
                            <div class="tab-content">
                                <div class="tab-pane active show" id="display">
                                    <div class="card-body">
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">显示用户审计记录</label>
                                            <div class="col">
                                                <select id="display_detect_log" class="col form-select"
                                                        value="{{index .settings "display_detect_log"}}">
                                                    <option value="0"
                                                            {{if not (.settings.display_detect_log)}}selected{{end}}>关闭
                                                    </option>
                                                    <option value="1" {{if .settings.display_detect_log}}selected{{end}}>
                                                        开启
                                                    </option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">显示文档</label>
                                            <div class="col">
                                                <select id="display_docs" class="col form-select"
                                                        value="{{index .settings "display_docs"}}">
                                                    <option value="0" {{if not (index .settings "display_docs")}}selected{{end}}>
                                                        关闭
                                                    </option>
                                                    <option value="1" {{if (index .settings "display_docs")}}selected{{end}}>开启
                                                    </option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">文档仅付费用户可见</label>
                                            <div class="col">
                                                <select id="display_docs_only_for_paid_user" class="col form-select"
                                                        value="{{index .settings "display_docs_only_for_paid_user"}}">
                                                    <option value="0"
                                                            {{if not (index .settings "display_docs_only_for_paid_user")}}selected{{end}}>
                                                        关闭
                                                    </option>
                                                    <option value="1"
                                                            {{if (index .settings "display_docs_only_for_paid_user")}}selected{{end}}>
                                                        开启
                                                    </option>
                                                </select>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                <div class="tab-pane" id="log">
                                    <div class="card-body">
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">启用每小时使用流量日志</label>
                                            <div class="col">
                                                <select id="traffic_log" class="col form-select"
                                                        value="{{index .settings "traffic_log"}}">
                                                    <option value="0" {{if not (index .settings "traffic_log")}}selected{{end}}>
                                                        关闭
                                                    </option>
                                                    <option value="1" {{if (index .settings "traffic_log")}}selected{{end}}>开启
                                                    </option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">流量日志保留天数</label>
                                            <div class="col">
                                                <input id="traffic_log_retention_days" type="text" class="form-control"
                                                       value="{{index .settings "traffic_log_retention_days"}}">
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">启用订阅日志</label>
                                            <div class="col">
                                                <select id="subscribe_log" class="col form-select"
                                                        value="{{index .settings "subscribe_log"}}">
                                                    <option value="0" {{if not (index .settings "subscribe_log")}}selected{{end}}>
                                                        关闭
                                                    </option>
                                                    <option value="1" {{if (index .settings "subscribe_log")}}selected{{end}}>
                                                        开启
                                                    </option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">订阅日志保留天数</label>
                                            <div class="col">
                                                <input id="subscribe_log_retention_days" type="text"
                                                       class="form-control"
                                                       value="{{index .settings "subscribe_log_retention_days"}}">
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">通知用户新IP订阅</label>
                                            <div class="col">
                                                <select id="notify_new_subscribe" class="col form-select"
                                                        value="{{index .settings "notify_new_subscribe"}}">
                                                    <option value="0"
                                                            {{if not (index .settings "notify_new_subscribe")}}selected{{end}}>关闭
                                                    </option>
                                                    <option value="1"
                                                            {{if (index .settings "notify_new_subscribe")}}selected{{end}}>开启
                                                    </option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">启用登录日志</label>
                                            <div class="col">
                                                <select id="login_log" class="col form-select"
                                                        value="{{index .settings "login_log"}}">
                                                    <option value="0" {{if not (index .settings "login_log")}}selected{{end}}>关闭
                                                    </option>
                                                    <option value="1" {{if (index .settings "login_log")}}selected{{end}}>开启
                                                    </option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">通知用户新IP登录</label>
                                            <div class="col">
                                                <select id="notify_new_login" class="col form-select"
                                                        value="{{index .settings "notify_new_login"}}">
                                                    <option value="0" {{if not (index .settings "notify_new_login")}}selected{{end}}>
                                                        关闭
                                                    </option>
                                                    <option value="1" {{if (index .settings "notify_new_login")}}selected{{end}}>
                                                        开启
                                                    </option>
                                                </select>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                <div class="tab-pane" id="checkin">
                                    <div class="card-body">
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">启用签到</label>
                                            <div class="col">
                                                <select id="enable_checkin" class="col form-select"
                                                        value="{{index .settings "enable_checkin"}}">
                                                    <option value="0" {{if not (index .settings "enable_checkin")}}selected{{end}}>
                                                        关闭
                                                    </option>
                                                    <option value="1" {{if (index .settings "enable_checkin")}}selected{{end}}>开启
                                                    </option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">签到最少流量（MB）</label>
                                            <div class="col">
                                                <input id="checkin_min" type="text" class="form-control"
                                                       value="{{index .settings "checkin_min"}}">
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">签到最多流量（MB）</label>
                                            <div class="col">
                                                <input id="checkin_max" type="text"
                                                       class="form-control"
                                                       value="{{index .settings "checkin_max"}}">
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <script>
            $("#save-setting").click(function () {
                $.ajax({
                    url: '/admin/setting/feature',
                    type: 'POST',
                    dataType: "json",
                    data: {
                        {{range $key := .update_field}}
                        {{$key}}: $('#{{$key}}').val(),
                        {{end}}
                    },
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
        </script>

        {{template "views/tabler/admin/footer.tpl" .}}
