{{template "views/tabler/admin/header.tpl" .}}

<div class="page-wrapper">
    <div class="container-xl">
        <div class="page-header d-print-none text-white">
            <div class="row align-items-center">
                <div class="col">
                    <h2 class="page-title">
                        <span class="home-title">定时任务设置</span>
                    </h2>
                    <div class="page-pretitle my-3">
                        <span class="home-subtitle">设置站点的定时任务</span>
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
                                    <a href="#daily_job" class="nav-link active" data-bs-toggle="tab">每日任务</a>
                                </li>
                                <li class="nav-item">
                                    <a href="#finance_mail" class="nav-link" data-bs-toggle="tab">财务报告</a>
                                </li>
                                <li class="nav-item">
                                    <a href="#detect" class="nav-link" data-bs-toggle="tab">审计任务</a>
                                </li>
                                <li class="nav-item">
                                    <a href="#inactive" class="nav-link" data-bs-toggle="tab">闲置账号检测</a>
                                </li>
                            </ul>
                        </div>
                        <div class="card-body">
                            <div class="tab-content">
                                <div class="tab-pane active show" id="daily_job">
                                    <div class="card-body">
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">每日任务执行时间(小时)</label>
                                            <div class="col">
                                                <input id="daily_job_hour" type="text" class="form-control"
                                                       value="{{index .settings "daily_job_hour"}}">
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">每日任务执行时间(分钟)</label>
                                            <div class="col">
                                                <input id="daily_job_minute" type="text" class="form-control"
                                                       value="{{index .settings "daily_job_minute"}}">
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                <div class="tab-pane show" id="finance_mail">
                                    <div class="card-body">
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">是否启用每日财务报告</label>
                                            <div class="col">
                                                <select id="enable_daily_finance_mail" class="col form-select"
                                                        value="{{index .settings "enable_daily_finance_mail"}}">
                                                    <option value="0"
                                                            {{if not (index .settings "enable_daily_finance_mail")}}selected{{end}}>
                                                        关闭
                                                    </option>
                                                    <option value="1"
                                                            {{if (index .settings "enable_daily_finance_mail")}}selected{{end}}>开启
                                                    </option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">是否启用每周财务报告</label>
                                            <div class="col">
                                                <select id="enable_weekly_finance_mail" class="col form-select"
                                                        value="{{index .settings "enable_weekly_finance_mail"}}">
                                                    <option value="0"
                                                            {{if not (index .settings "enable_weekly_finance_mail")}}selected{{end}}>
                                                        关闭
                                                    </option>
                                                    <option value="1"
                                                            {{if (index .settings "enable_weekly_finance_mail")}}selected{{end}}>开启
                                                    </option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">是否启用每月财务报告</label>
                                            <div class="col">
                                                <select id="enable_monthly_finance_mail" class="col form-select"
                                                        value="{{index .settings "enable_monthly_finance_mail"}}">
                                                    <option value="0"
                                                            {{if not (index .settings "enable_monthly_finance_mail")}}selected{{end}}>
                                                        关闭
                                                    </option>
                                                    <option value="1"
                                                            {{if (index .settings "enable_monthly_finance_mail")}}selected{{end}}>
                                                        开启
                                                    </option>
                                                </select>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                <div class="tab-pane show" id="detect">
                                    <div class="card-body">
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">是否启用节点被墙检测</label>
                                            <div class="col">
                                                <select id="enable_detect_gfw" class="col form-select"
                                                        value="{{index .settings "enable_detect_gfw"}}">
                                                    <option value="0"
                                                            {{if not (index .settings "enable_detect_gfw")}}selected{{end}}>关闭
                                                    </option>
                                                    <option value="1" {{if (index .settings "enable_detect_gfw")}}selected{{end}}>
                                                        开启
                                                    </option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">是否启用审计封禁</label>
                                            <div class="col">
                                                <select id="enable_detect_ban" class="col form-select"
                                                        value="{{index .settings "enable_detect_ban"}}">
                                                    <option value="0"
                                                            {{if not (index .settings "enable_detect_ban")}}selected{{end}}>关闭
                                                    </option>
                                                    <option value="1" {{if (index .settings "enable_detect_ban")}}selected{{end}}>
                                                        开启
                                                    </option>
                                                </select>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                <div class="tab-pane show" id="inactive">
                                    <div class="card-body">
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">是否启用闲置账号检测</label>
                                            <div class="col">
                                                <select id="enable_detect_inactive_user" class="col form-select"
                                                        value="{{index .settings "enable_detect_inactive_user"}}">
                                                    <option value="0"
                                                            {{if not (index .settings "enable_detect_inactive_user")}}selected{{end}}>
                                                        关闭
                                                    </option>
                                                    <option value="1"
                                                            {{if (index .settings "enable_detect_inactive_user")}}selected{{end}}>
                                                        开启
                                                    </option>
                                                </select>
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">未签到时长(天)</label>
                                            <div class="col">
                                                <input id="detect_inactive_user_checkin_days" type="text"
                                                       class="form-control"
                                                       value="{{index .settings "detect_inactive_user_checkin_days"}}">
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">未登录时长(天)</label>
                                            <div class="col">
                                                <input id="detect_inactive_user_login_days" type="text"
                                                       class="form-control"
                                                       value="{{index .settings "detect_inactive_user_login_days"}}">
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">未使用时长(天)</label>
                                            <div class="col">
                                                <input id="detect_inactive_user_use_days" type="text"
                                                       class="form-control"
                                                       value="{{index .settings "detect_inactive_user_use_days"}}">
                                            </div>
                                        </div>
                                        <div class="form-group mb-3 row">
                                            <label class="form-label col-3 col-form-label">是否启用移除闲置账号订阅链接与邀请码</label>
                                            <div class="col">
                                                <select id="remove_inactive_user_link_and_invite" class="col form-select"
                                                        value="{{.settings.remove_inactive_user_link_and_invite}}">
                                                    <option value="0"
                                                            {{if not .settings.remove_inactive_user_link_and_invite}}selected{{end}}>
                                                        关闭
                                                    </option>
                                                    <option value="1"
                                                            {{if .settings.remove_inactive_user_link_and_invite}}selected{{end}}>
                                                        开启
                                                    </option>
                                                </select>
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
                    url: '/admin/setting/cron',
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
