{{template "views/tabler/user/header.tpl" .}}

<script src="//{{index .config "jsdelivr_url"}}/npm/jquery/dist/jquery.min.js"></script>

<div class="page-wrapper">
    <div class="container-xl">
        <div class="page-header d-print-none text-white">
            <div class="row align-items-center">
                <div class="col">
                    <h2 class="page-title">
                        <span class="home-title my-3">账单 #{{.invoice.Id}}</span>
                    </h2>
                    <div class="page-pretitle">
                        <span class="home-subtitle">账单详情</span>
                    </div>
                </div>
                <div class="col-auto ms-auto d-print-none">
                    <div class="btn-list">
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="page-body">
        <div class="container-xl">
            <div class="row row-cards">
                {{if eq .invoice.Status "unpaid"}}
                <div class="col-sm-12 col-md-6 col-lg-9">
                    {{else}}
                    <div class="col-sm-12 col-md-12 col-lg-12">
                        {{end}}
                        <div class="card">
                            <div class="card-header">
                                <h3 class="card-title">基本信息</h3>
                            </div>
                            <div class="card-body">
                                <div class="datagrid">
                                    <div class="datagrid-item">
                                        <div class="datagrid-title">订单ID</div>
                                        <div class="datagrid-content">{{.invoice.OrderId}}</div>
                                    </div>
                                    <div class="datagrid-item">
                                        <div class="datagrid-title">订单金额</div>
                                        <div class="datagrid-content">{{.invoice.Price}}</div>
                                    </div>
                                    <div class="datagrid-item">
                                        <div class="datagrid-title">订单状态</div>
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
                                    {{if eq .invoice.Status "paid_gateway"}}
                                        <div class="datagrid-item">
                                            <div class="datagrid-title">支付网关单号</div>
                                            <div class="datagrid-content">{{.paylist.Tradeno}}</div>
                                        </div>
                                    {{end}}
                                </div>
                            </div>
                        </div>
                        <div class="card my-3">
                            <div class="card-header">
                                <h3 class="card-title">账单详情</h3>
                            </div>
                            <div class="card-body">
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
                                                <td>{{.invoice_content.name}}</td>
                                                <td>{{.invoice_content.price}}</td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </div>
                            </div>
                        </div>
                    </div>
                    {{if eq .invoice.Status "unpaid"}}
                        <div class="col-sm-12 col-md-6 col-lg-3">
                            <div class="card">
                                <div class="card-header">
                                    <h3 class="card-title">余额支付</h3>
                                </div>
                                <div class="card-body">
                                    <div class="mb-3">
                                        当前账户可用余额：<code>{{.user.Money}}</code> 元
                                    </div>
                                </div>
                                <div class="card-footer">
                                    <div class="d-flex">
                                        <button id="pay-balance" class="btn btn-blue" type="button">支付</button>
                                    </div>
                                </div>
                            </div>
                            {{if gt (.payments|len) 0}}
                                <div class="card my-3">
                                    <div class="card-header">
                                        <h3 class="card-title">网关支付</h3>
                                    </div>
                                    <div class="card-body">
                                        {{range $payment := .payments}}
                                            <div class="mb-3">
                                                {{if eq $payment.Name "f2f"}}
                                                    {{template "views/tabler/gateway/f2f.tpl" $}}
                                                {{else if eq $payment.Name "stripe"}}
                                                    {{template "views/tabler/gateway/stripe.tpl" $}}
                                                {{else if eq $payment.Name "epay"}}
                                                    {{template "views/tabler/gateway/epay.tpl" $}}
                                                {{else}}
                                                    {{template "views/tabler/gateway/paypal.tpl" $}}
                                                {{end}}
                                            </div>
                                        {{end}}
                                    </div>
                                </div>
                            {{end}}
                        </div>
                    {{end}}
                </div>
            </div>
        </div>

        <script>
            $("#pay-balance").click(function () {
                $.ajax({
                    url: '/user/invoice/pay_balance',
                    type: 'POST',
                    dataType: "json",
                    data: {
                        invoice_id: {{.invoice.Id}},
                    },
                    success: function (data) {
                        if (data.ret == 1) {
                            $('#success-message').text(data.msg);
                            $('#success-dialog').modal('show');
                            setTimeout(function () {
                                $(location).attr('href', '/user/invoice');
                            }, {{index .config "jump_delay"}});
                        } else {
                            $('#fail-message').text(data.msg);
                            $('#fail-dialog').modal('show');
                        }
                    }
                })
            });
        </script>

        {{template "views/tabler/user/footer.tpl" .}}
