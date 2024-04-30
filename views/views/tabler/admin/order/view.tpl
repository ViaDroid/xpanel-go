{{template "views/tabler/admin/header.tpl" .}}

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
                        <a href="/admin/user/{{.order.UserId}}/edit" targer="_blank" class="btn btn-primary">
                            <i class="icon ti ti-user"></i>
                            查看关联用户
                        </a>
                        <a href="/admin/invoice/{{.invoice.Id}}/view" targer="_blank" class="btn btn-primary">
                            <i class="icon ti ti-file-dollar"></i>
                            查看关联账单
                        </a>
                        {{if eq "pending_payment" .order.Status}}
                            <button href="#" class="btn btn-red" data-bs-toggle="modal"
                                    data-bs-target="#cancel_order_confirm_dialog">
                                <i class="icon ti ti-x"></i>
                                取消订单
                            </button>
                        {{end}}
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
                            <div class="datagrid-title">提交用户</div>
                            <div class="datagrid-content">{{.order.UserId}}</div>
                        </div>
                        <div class="datagrid-item">
                            <div class="datagrid-title">商品ID</div>
                            <div class="datagrid-content">{{.order.ProductId}}</div>
                        </div>
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
                        {{if or (eq "tabp" .order.ProductType) (eq "time" .order.ProductType)}}
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
                        {{if or (eq "tabp" .order.ProductType) (eq "bandwidth" .order.ProductType)}}
                            <div class="datagrid-item">
                                <div class="datagrid-title">可用流量 (GB)</div>
                                <div class="datagrid-content">{{.order.ContentMap.bandwidth}}</div>
                            </div>
                        {{end}}
                        {{if or (eq "tabp" .order.ProductType) (eq "time" .order.ProductType)}}
                            <div class="datagrid-item">
                                <div class="datagrid-title">用户分组</div>
                                <div class="datagrid-content">{{.order.ContentMap.node_group}}</div>
                            </div>
                            <div class="datagrid-item">
                                <div class="datagrid-title">速率限制 (Mbps)</div>
                                <div class="datagrid-content">
                                    {{if czeq (.order.ContentMap.ip_limit) 0}}
                                        不限制
                                    {{else}}
                                        {{.order.ContentMap.speed_limit}}
                                    {{end}}
                                </div>
                            </div>
                            <div class="datagrid-item">
                                <div class="datagrid-title">同时连接IP限制</div>
                                <div class="datagrid-content">
                                    {{if czeq (.order.ContentMap.ip_limit) 0}}
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
                                         <td>{{.invoice.ContentMap.name}}</td>
                                         <td>{{.invoice.ContentMap.price}}</td>
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

    <div class="modal modal-blur fade" id="cancel_order_confirm_dialog" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered modal-dialog-scrollable" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">取消订单</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <div class="mb-3">
                        <p>
                            确认取消此订单？
                        <p>
                    </div>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn me-auto" data-bs-dismiss="modal">取消</button>
                    <button id="confirm_cancel" type="button" class="btn btn-primary" data-bs-dismiss="modal">确认
                    </button>
                </div>
            </div>
        </div>
    </div>

    <script>
        $("#confirm_cancel").click(function () {
            $.ajax({
                url: "/admin/order/{{.order.Id}}/cancel",
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
    </script>

    {{template "views/tabler/admin/footer.tpl" .}}
