<div class="card-inner">
    <h4>
        EPay 在线充值
    </h4>
    <p class="card-heading"></p>
    <form class="epay" name="epay" action="/user/payment/purchase/epay" method="post">
        <input hidden id="price" name="price" value="{{.invoice.Price}}">
        <input hidden id="invoice_id" name="invoice_id" value="{{.invoice.Id}}">
        {{if (index .public_setting "epay_alipay")}}
            <button class="btn btn-flat waves-attach" id="btnSubmit" type="submit" name="type" value="alipay">
                <img src="/images/alipay.png" height="50px"/>
            </button>
        {{end}}
        {{if (index .public_setting "epay_wechat")}}
            <button class="btn btn-flat waves-attach" id="btnSubmit" type="submit" name="type" value="wxpay">
                <img src="/images/wechat.png" height="50px"/>
            </button>
        {{end}}
        {{if (index .public_setting "epay_qq")}}
            <button class="btn btn-flat waves-attach" id="btnSubmit" type="submit" name="type" value="qqpay">
                <img src="/images/qqpay.png" height="50px"/>
            </button>
        {{end}}
        {{if (index .public_setting "epay_usdt")}}
            <button class="btn btn-flat waves-attach" id="btnSubmit" type="submit" name="type" value="usdt">
                <img src="/images/usdt.png" height="50px"/>
            </button>
        {{end}}
    </form>
</div>
