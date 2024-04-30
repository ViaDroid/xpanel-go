{{template "views/tabler/admin/header.tpl" .}}

<div class="page-wrapper">
    <div class="container-xl">
        <div class="page-header d-print-none text-white">
            <div class="row align-items-center">
                <div class="col">
                    <h2 class="page-title">
                        <span class="home-title">用户 #{{.edit_user.Id}}</span>
                    </h2>
                    <div class="page-pretitle my-3">
                        <span class="home-subtitle">用户编辑</span>
                    </div>
                </div>
                <div class="col-auto">
                    <div class="btn-list">
                        <a id="save_changes" href="#" class="btn btn-primary">
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
                <div class="col-md-4 col-sm-12">
                    <div class="card">
                        <div class="card-header card-header-light">
                            <h3 class="card-title">基础信息</h3>
                        </div>
                        <div class="card-body">
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">注册邮箱</label>
                                <div class="col">
                                    <input id="email" type="email" class="form-control" value="{{.edit_user.Email}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">用户名</label>
                                <div class="col">
                                    <input id="user_name" type="text" class="form-control"
                                           value="{{.edit_user.UserName}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">备注</label>
                                <div class="col">
                                    <input id="remark" type="text" class="form-control" value="{{.edit_user.Remark}}"
                                           placeholder="仅管理员可见">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">账户密码</label>
                                <div class="col">
                                    <input id="pass" type="text" class="form-control"
                                           placeholder="若需为此用户重置密码, 填写此栏">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">账户余额</label>
                                <div class="col">
                                    <input id="money" type="number" step="0.1" class="form-control"
                                           value="{{.edit_user.Money}}">
                                </div>
                            </div>
                            <div class="hr-text">
                                <span>时间设置</span>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-4 col-form-label">等级过期时间</label>
                                <div class="col">
                                    <input id="class_expire" type="text" class="form-control"
                                           value="{{.edit_user.ClassExpire}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-4 col-form-label">免费用户流量重置日</label>
                                <div class="col">
                                    <input id="auto_reset_day" type="text" class="form-control"
                                           value="{{.edit_user.AutoResetDay}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-4 col-form-label">
                                    重置的免费流量(GB)
                                </label>
                                <div class="col">
                                    <input id="auto_reset_bandwidth" type="text" class="form-control"
                                           value="{{.edit_user.AutoResetBandwidth}}">
                                </div>
                            </div>
                            <div class="hr-text">
                                <span>高级选项</span>
                            </div>
                            <div class="form-group mb-3 row">
                                <span class="col">管理员</span>
                                <span class="col-auto">
                                    <label class="form-check form-check-single form-switch">
                                        <input id="is_admin" class="form-check-input" type="checkbox"
                                               {{if .edit_user.IsAdmin}}checked="" {{end}}>
                                    </label>
                                </span>
                            </div>
                            <div class="form-group mb-3 row">
                                <span class="col">两步认证</span>
                                <span class="col-auto">
                                    <label class="form-check form-check-single form-switch">
                                        <input id="ga_enable" class="form-check-input" type="checkbox"
                                               {{if .edit_user.GaEnable}}checked="" {{end}}>
                                    </label>
                                </span>
                            </div>
                            <div class="form-group mb-3 row">
                                <span class="col">账户异常状态</span>
                                <span class="col-auto form-check-single form-switch">
                                    <input id="is_shadow_banned" class="form-check-input" type="checkbox"
                                           {{if .edit_user.IsShadowBanned}}checked=""{{end}}>
                                </span>
                            </div>
                            <div class="form-group mb-3 row">
                                <span class="col">封禁用户</span>
                                <span class="col-auto">
                                    <label class="form-check form-check-single form-switch">
                                        <input id="is_banned" class="form-check-input" type="checkbox"
                                               {{if .edit_user.IsBanned}}checked=""{{end}}>
                                    </label>
                                </span>
                            </div>
                            <div class="form-group mb-3 row">
                                <span class="col">手动封禁理由</span>
                                <span class="col-auto">
                                    <input id="banned_reason" type="text" class="form-control"
                                           value="{{.edit_user.BannedReason}}">
                                </span>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-md-4 col-sm-12">
                    <div class="card">
                        <div class="card-header card-header-light">
                            <h3 class="card-title">其他信息</h3>
                        </div>
                        <div class="card-body">
                            <div class="form-group mb-3 row">
                                <label class="form-label col-4 col-form-label">流量限制</label>
                                <div class="col">
                                    <input id="transfer_enable" type="text" class="form-control"
                                           value="{{.edit_user.EnableTraffic}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-4 col-form-label">当期用量</label>
                                <div class="col">
                                    <input id="usedTraffic" type="text" class="form-control"
                                           value="{{.edit_user.UsedTraffic}}" disabled/>
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-4 col-form-label">累计用量</label>
                                <div class="col">
                                    <input id="usedTraffic" type="text" class="form-control"
                                           value="{{.edit_user.TotalTraffic}}" disabled/>
                                </div>
                            </div>
                            <div class="hr-text">
                                <span>邀请注册</span>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-4 col-form-label">邀请人</label>
                                <div class="col">
                                    <input id="ref_by" type="text" class="form-control" value="{{.edit_user.RefBy}}">
                                </div>
                            </div>
                            <div class="hr-text">
                                <span>划分与使用限制</span>
                            </div>
                            <div class="form-group mb-3 col-12">
                                <label class="form-label col-12 col-form-label">节点群组</label>
                                <div class="col">
                                    <input id="node_group" type="text" class="form-control"
                                           value="{{.edit_user.NodeGroup}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 col-12">
                                <label class="form-label col-12 col-form-label">账户等级</label>
                                <div class="col">
                                    <input id="class" type="text" class="form-control" value="{{.edit_user.Class}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 col-12">
                                <label class="form-label col-12 col-form-label">速度限制 (Mbps)</label>
                                <div class="col">
                                    <input id="node_speedlimit" type="text" class="form-control"
                                           value="{{.edit_user.NodeSpeedlimit}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 col-12">
                                <label class="form-label col-12 col-form-label">同時连接 IP 限制</label>
                                <div class="col">
                                    <input id="node_iplimit" type="text" class="form-control"
                                           value="{{.edit_user.NodeIplimit}}">
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="col-md-4 col-sm-12">
                    <div class="card">
                        <div class="card-header card-header-light">
                            <h3 class="card-title">连接设置</h3>
                        </div>
                        <div class="card-body">
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">端口</label>
                                <div class="col">
                                    <input id="port" type="text" class="form-control" value="{{.edit_user.Port}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">密码</label>
                                <div class="col">
                                    <input id="passwd" type="text" class="form-control" value="{{.edit_user.Passwd}}">
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">加密</label>
                                <div class="col">
                                    <input id="method" type="text" class="form-control" value="{{.edit_user.Method}}">
                                </div>
                            </div>
                            <div class="hr-text">
                                <span>访问限制</span>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">IP / CIDR</label>
                                <div class="col">
                                    <textarea id="forbidden_ip" class="col form-control"
                                              rows="2">{{.edit_user.ForbiddenIp}}</textarea>
                                </div>
                            </div>
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">PORT</label>
                                <div class="col">
                                    <textarea id="forbidden_port" class="col form-control"
                                              rows="2">{{.edit_user.ForbiddenPort}}</textarea>
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
    $("#save_changes").click(function () {
        $.ajax({
            url: '/admin/user/{{.edit_user.Id}}',
            type: 'PUT',
            dataType: "json",
            data: {
                {{range $key := .update_field}}
                {{$key}}: $('#{{$key}}').val(),
                {{end}}
                is_admin: $("#is_admin").is(":checked"),
                is_banned: $("#is_banned").is(":checked"),
                ga_enable: $("#ga_enable").is(":checked"),
                is_shadow_banned: $("#is_shadow_banned").is(":checked"),
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
