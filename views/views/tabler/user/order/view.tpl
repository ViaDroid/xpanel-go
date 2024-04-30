{{template "views/tabler/user/header.tpl" .}}

<div class="page-wrapper">
    <div class="container-xl">
        <div class="page-header d-print-none text-white">
            <div class="row align-items-center">
                <div class="col">
                    <h2 class="page-title">
                        <span class="home-title">订单 #{{.order.Id}}</span>
                    </h2>
                    <div class="page-pretitle my-3">
                        <span class="home-subtitle">订单详情</span>
                    </div>
                </div>
                <div class="col-auto">
                    <div class="btn-list">
                        <a href="/user/invoice/{{.invoice.Id}}/view" targer="_blank" class="btn btn-primary">
                            <i class="icon ti ti-file-dollar"></i>
                            查看账单
                        </a>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="page-body">
        <div class="container-xl">
            <div class="card">
                <div class="card-header">
                    <h3 class="card-title">基本信息</h3>
                </div>
                <div class="card-body">
                    <div class="datagrid">
                        <div class="datagrid-item">
                            <div class="datagrid-title">商品类型</div>
                            <div class="datagrid-content">{{.order.ProductTypeStr}}</div>
                        </div>
                        <div class="datagrid-item">
                            <div class="datagrid-title">商品名称</div>
                            <div class="datagrid-content">{{.order.ProductName}}</div>
                        </div>
                        <div class="datagrid-item">
                            <div class="datagrid-title">订单优惠码</div>
                            <div class="datagrid-content">{{.order.Coupon}}</div>
                        </div>
                        <div class="datagrid-item">
                            <div class="datagrid-title">订单金额</div>
                            <div class="datagrid-content">{{.order.Price}}</div>
                        </div>
                        <div class="datagrid-item">
                            <div class="datagrid-title">订单状态</div>
                            <div class="datagrid-content">{{.order.StatusStr}}</div>
                        </div>
                        <div class="datagrid-item">
                            <div class="datagrid-title">创建时间</div>
                            <div class="datagrid-content">{{.order.CreateTimeStr}}</div>
                        </div>
                        <div class="datagrid-item">
                            <div class="datagrid-title">更新时间</div>
                            <div class="datagrid-content">{{.order.UpdateTimeStr}}</div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="card my-3">
                <div class="card-header">
                    <h3 class="card-title">商品内容</h3>
                </div>
                <div class="card-body">
                    <div class="datagrid">
                        {{if or (eq .order.ProductType "tabp") (eq .order.ProductType "time")}}
                            <div class="datagrid-item">
                                <div class="datagrid-title">商品时长 (天)</div>
                                <div class="datagrid-content">{{.order.ContentMap.time}}</div>
                            </div>
                            <div class="datagrid-item">
                                <div class="datagrid-title">等级时长 (天)</div>
                                <div class="datagrid-content">{{.order.ContentMap.class_time}}</div>
                            </div>
                            <div class="datagrid-item">
                                <div class="datagrid-title">等级</div>
                                <div class="datagrid-content">{{.order.ContentMap.class}}</div>
                            </div>
                        {{end}}
                        {{if or (eq .order.ProductType "tabp") (eq .order.product_type "bandwidth")}}
                            <div class="datagrid-item">
                                <div class="datagrid-title">可用流量 (GB)</div>
                                <div class="datagrid-content">{{.order.ContentMap.bandwidth}}</div>
                            </div>
                        {{end}}
                        {{if or (eq .order.ProductType "tabp") (eq .order.ProductType "time")}}
                            <div class="datagrid-item">
                                <div class="datagrid-title">速率限制 (Mbps)</div>
                                <div class="datagrid-content">
                                    {{if czeq .order.ContentMap.ip_limit 0}}
                                        不限制
                                    {{else}}
                                        {{.order.ContentMap.speed_limit}}
                                    {{end}}
                                </div>
                            </div>
                            <div class="datagrid-item">
                                <div class="datagrid-title">同时连接IP限制</div>
                                <div class="datagrid-content">
                                    {{if czeq .order.ContentMap.ip_limit 0}}
                                        不限制
                                    {{else}}
                                        {{.order.ContentMap.ip_limit}}
                                    {{end}}
                                </div>
                            </div>
                        {{end}}
                    </div>
                </div>
            </div>
            <div class="card my-3">
                <div class="card-header">
                    <h3 class="card-title">关联账单</h3>
                </div>
                <div class="card-body">
                    <div class="datagrid">
                        <div class="datagrid-item">
                            <div class="datagrid-title">账单内容</div>
                            <div class="datagrid-content">
                                <div class="table-responsive">
                                    <table id="invoice_content_table" class="table table-vcenter card-table">
                                        <thead>
                                        <tr>
                                            <th>名称</th>
                                            <th>价格</th>
                                        </tr>
                                        </thead>
                                        <tbody>
                                            <tr>
                                                <td>{{.invoice.ContentMap.name}}</td>
                                                <td>{{.invoice.ContentMap.price}}</td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </div>
                            </div>
                        </div>
                        <div class="datagrid-item">
                            <div class="datagrid-title">账单金额</div>
                            <div class="datagrid-content">{{.invoice.Price}}</div>
                        </div>
                        <div class="datagrid-item">
                            <div class="datagrid-title">账单状态</div>
                            <div class="datagrid-content">{{.invoice.StatusStr}}</div>
                        </div>
                        <div class="datagrid-item">
                            <div class="datagrid-title">创建时间</div>
                            <div class="datagrid-content">{{.invoice.CreateTimeStr}}</div>
                        </div>
                        <div class="datagrid-item">
                            <div class="datagrid-title">更新时间</div>
                            <div class="datagrid-content">{{.invoice.UpdateTimeStr}}</div>
                        </div>
                        <div class="datagrid-item">
                            <div class="datagrid-title">支付时间</div>
                            <div class="datagrid-content">{{.invoice.PayTimeStr}}</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    {{template "views/tabler/user/footer.tpl" .}}
