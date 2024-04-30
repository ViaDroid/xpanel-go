{{template "views/tabler/user/header.tpl" .}}

<div class="page-wrapper">
    <div class="container-xl">
        <div class="page-header d-print-none text-white">
            <div class="row align-items-center">
                <div class="col">
                    <h2 class="page-title">
                        <span class="home-title">商品列表</span>
                    </h2>
                    <div class="page-pretitle my-3">
                        <span class="home-subtitle">浏览你所需要的商品</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="page-body">
        <div class="container-xl">
            <div class="row row-cards">
                <div class="col-12">
                    <div class="card">
                        <ul class="nav nav-tabs nav-fill" data-bs-toggle="tabs">
                            <li class="nav-item">
                                <a href="#tabp" class="nav-link active" data-bs-toggle="tab">
                                    <i class="ti ti-rotate-360 icon"></i>
                                    &nbsp;时间流量包
                                </a>
                            </li>
                            <li class="nav-item">
                                <a href="#bandwidth" class="nav-link" data-bs-toggle="tab">
                                    <i class="ti ti-arrows-down-up icon"></i>
                                    &nbsp;流量包
                                </a>
                            </li>
                            <li class="nav-item">
                                <a href="#time" class="nav-link" data-bs-toggle="tab">
                                    <i class="ti ti-clock icon"></i>
                                    &nbsp;时间包
                                </a>
                            </li>
                        </ul>
                        <div class="card-body">
                            <div class="tab-content">
                                <div class="tab-pane active show" id="tabp">
                                    <div class="row">
                                        {{range  $tabp := .tabps}}
                                            <div class="col-md-3 col-sm-12 my-3">
                                                <div class="card card-md">
                                                    <div class="card-body text-center">
                                                        <div id="product-{{$tabp.Id}}-name"
                                                             class="text-uppercase text-secondary font-weight-medium">
                                                            {{$tabp.Name}}</div>
                                                        <div id="product-{{$tabp.Id}}-price"
                                                             class="display-6 my-3">
                                                            <p class="fw-bold">{{$tabp.Price}}</p>
                                                            <i class="ti ti-currency-yuan"></i>
                                                        </div>
                                                        <div class="list-group list-group-flush">
                                                            <div class="list-group-item">
                                                                <div class="row align-items-center">
                                                                    <div class="col text-truncate">
                                                                        <div class="text-reset d-block">
                                                                            Lv. {{$tabp.ContentMap.class}}</div>
                                                                        <div class="d-block text-secondary text-truncate mt-n1">
                                                                            等级
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                            </div>
                                                            <div class="list-group-item">
                                                                <div class="row align-items-center">
                                                                    <div class="col text-truncate">
                                                                        <div class="text-reset d-block">{{$tabp.ContentMap.class_time}}
                                                                            天
                                                                        </div>
                                                                        <div class="d-block text-secondary text-truncate mt-n1">
                                                                            等级时长
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                            </div>
                                                            <div class="list-group-item">
                                                                <div class="row align-items-center">
                                                                    <div class="col text-truncate">
                                                                        <div class="text-reset d-block">{{$tabp.ContentMap.bandwidth}}
                                                                            GB
                                                                        </div>
                                                                        <div class="d-block text-secondary text-truncate mt-n1">
                                                                            可用流量
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                            </div>
                                                            <div class="list-group-item">
                                                                <div class="row align-items-center">
                                                                    <div class="col text-truncate">
                                                                        {{if czeq $tabp.ContentMap.speed_limit 0}}
                                                                            <div class="text-reset d-block">不限制</div>
                                                                        {{else}}
                                                                            <div class="text-reset d-block">{{$tabp.ContentMap.speed_limit}}
                                                                                Mbps
                                                                            </div>
                                                                        {{end}}
                                                                        <div class="d-block text-secondary text-truncate mt-n1">
                                                                            连接速度
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                            </div>
                                                            <div class="list-group-item">
                                                                <div class="row align-items-center">
                                                                    <div class="col text-truncate">
                                                                        {{if czeq $tabp.ContentMap.ip_limit 0}}
                                                                            <div class="text-reset d-block">不限制</div>
                                                                        {{else}}
                                                                            <div class="text-reset d-block">{{$tabp.ContentMap.ip_limit}}</div>
                                                                        {{end}}
                                                                        <div class="d-block text-secondary text-truncate mt-n1">
                                                                            同时连接 IP 数
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                            </div>
                                                        </div>
                                                        <div class="row g-2">
                                                            {{if or (eq $tabp.Stock -1) (gt $tabp.Stock 0)}}
                                                                <div class="col">
                                                                    <a href="/user/order/create?product_id={{$tabp.Id}}"
                                                                       class="btn btn-primary w-100 my-3">购买</a>
                                                                </div>
                                                            {{else}}
                                                                <div class="col">
                                                                    <a href="" class="btn btn-primary w-100 my-3"
                                                                       disabled>告罄</a>
                                                                </div>
                                                            {{end}}
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        {{end}}
                                    </div>
                                </div>
                                <div class="tab-pane show" id="bandwidth">
                                    <div class="row">
                                        {{range  $bandwidth := .bandwidths}}
                                            <div class="col-md-3 col-sm-12 my-3">
                                                <div class="card card-md">
                                                    <div class="card-body text-center">
                                                        <div id="product-{{$bandwidth.Id}}-name"
                                                             class="text-uppercase text-secondary font-weight-medium">
                                                            {{$bandwidth.Name}}</div>
                                                        <div id="product-{{$bandwidth.Id}}-price"
                                                             class="display-6 my-3">
                                                            <p class="fw-bold">{{$bandwidth.Price}}</p>
                                                            <i class="ti ti-currency-yuan"></i>
                                                        </div>
                                                        <div class="list-group list-group-flush">
                                                            <div class="list-group-item">
                                                                <div class="row align-items-center">
                                                                    <div class="col text-truncate">
                                                                        <div class="text-reset d-block">{{$bandwidth.ContentMap.bandwidth}}
                                                                            GB
                                                                        </div>
                                                                        <div class="d-block text-secondary text-truncate mt-n1">
                                                                            可用流量
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                            </div>
                                                        </div>
                                                        <div class="row g-2">
                                                            {{if or (eq $bandwidth.Stock -1) (gt $bandwidth.Stock 0)}}
                                                                <div class="col">
                                                                    <a href="/user/order/create?product_id={{$bandwidth.Id}}"
                                                                       class="btn btn-primary w-100 my-3">购买</a>
                                                                </div>
                                                            {{else}}
                                                                <div class="col">
                                                                    <a href="" class="btn btn-primary w-100 my-3"
                                                                       disabled>告罄</a>
                                                                </div>
                                                            {{end}}
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        {{end}}
                                    </div>
                                </div>
                                <div class="tab-pane show" id="time">
                                    <div class="row">
                                        {{range  $time := .times}}
                                            <div class="col-md-3 col-sm-12 my-3">
                                                <div class="card card-md">
                                                    <div class="card-body text-center">
                                                        <div id="product-{{$time.Id}}-name"
                                                             class="text-uppercase text-secondary font-weight-medium">
                                                            {{$time.Name}}
                                                        </div>
                                                        <div id="product-{{$time.Id}}-price"
                                                             class="display-6 my-3"><p
                                                                    class="fw-bold">{{$time.Price}}</p>
                                                            <i class="ti ti-currency-yuan"></i>
                                                        </div>
                                                        <div class="list-group list-group-flush">
                                                            <div class="list-group-item">
                                                                <div class="row align-items-center">
                                                                    <div class="col text-truncate">
                                                                        <div class="text-reset d-block">
                                                                            Lv. {{$time.ContentMap.class}}</div>
                                                                        <div class="d-block text-secondary text-truncate mt-n1">
                                                                            等级
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                            </div>
                                                            <div class="list-group-item">
                                                                <div class="row align-items-center">
                                                                    <div class="col text-truncate">
                                                                        <div class="text-reset d-block">{{$time.ContentMap.class_time}}
                                                                            天
                                                                        </div>
                                                                        <div class="d-block text-secondary text-truncate mt-n1">
                                                                            等级时长
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                            </div>
                                                            <div class="list-group-item">
                                                                <div class="row align-items-center">
                                                                    <div class="col text-truncate">
                                                                        {{if czeq $time.ContentMap.speed_limit 0}}
                                                                            <div class="text-reset d-block">不限制</div>
                                                                        {{else}}
                                                                            <div class="text-reset d-block">{{$time.ContentMap.speed_limit}}
                                                                                Mbps
                                                                            </div>
                                                                        {{end}}
                                                                        <div class="d-block text-secondary text-truncate mt-n1">
                                                                            连接速度
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                            </div>
                                                            <div class="list-group-item">
                                                                <div class="row align-items-center">
                                                                    <div class="col text-truncate">
                                                                        {{if czeq $time.ContentMap.ip_limit 0}}
                                                                            <div class="text-reset d-block">不限制</div>
                                                                        {{else}}
                                                                            <div class="text-reset d-block">{{$time.ContentMap.ip_limit}}</div>
                                                                        {{end}}
                                                                        <div class="d-block text-secondary text-truncate mt-n1">
                                                                            同时连接 IP 数
                                                                        </div>
                                                                    </div>
                                                                </div>
                                                            </div>
                                                        </div>
                                                        <div class="row g-2">
                                                            {{if or (eq $time.Stock -1) (gt $time.Stock 0)}}
                                                                <div class="col">
                                                                    <a href="/user/order/create?product_id={{$time.Id}}"
                                                                       class="btn btn-primary w-100 my-3">购买</a>
                                                                </div>
                                                            {{else}}
                                                                <div class="col">
                                                                    <a href="" class="btn btn-primary w-100 my-3"
                                                                       disabled>告罄</a>
                                                                </div>
                                                            {{end}}
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                        {{end}}
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    {{template "views/tabler/user/footer.tpl" .}}
