{{template "views/tabler/user/header.tpl" .}}

<div class="page-wrapper">
    <div class="container-xl">
        <div class="page-header d-print-none text-white">
            <div class="row align-items-center">
                <div class="col">
                    <h2 class="page-title">
                        <span class="home-title">账户已被封禁</span>
                    </h2>
                    <div class="page-pretitle my-3">
                        <span class="home-subtitle">你的账户功能已被停用，并且禁止访问用户中心</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="page-body">
        <div class="container-xl">
            <div class="row row-deck row-cards">
                <div class="col-sm-12 col-lg-12">
                    <div class="card">
                        <div class="empty">
                            <div class="empty-img">
                                <i class="ti ti-circle-x icon mb-2 text-danger icon-lg" style="font-size:3.5rem;"></i>
                            </div>
                            {{if eq .banned_reason "DetectBan"}}
                                <p class="empty-title">审计封禁</p>
                                <p class="empty-subtitle text-secondary">你的账户因为触发审计规则而被系统自动封禁</p>
                            {{else}}
                                <p class="empty-title">以下是你被封禁的理由</p>
                                <p class="empty-subtitle text-secondary">{$banned_reason}</p>
                            {{end}}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

{{template "views/tabler/user/footer.tpl" .}}