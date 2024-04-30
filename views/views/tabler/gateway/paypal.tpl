<script src="https://www.paypal.com/sdk/js?client-id={{index .public_setting "paypal_client_id"}}&currency={{index .public_setting "paypal_currency"}}"></script>

<div class="card-inner">
    <h4>
        PayPal
    </h4>
    <p class="card-heading"></p>
    <div id="paypal-button-container"></div>
</div>

<script>
    paypal.Buttons({
        createOrder() {
            return fetch("/user/payment/purchase/paypal", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    price: {{.invoice.Price}},
                    invoice_id: {{.invoice.Id}},
                }),
            })
                .then((response) => response.json())
                .then((order) => order.id);
        },
        onApprove(data) {
            return fetch("/payment/notify/paypal", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    order_id: data.orderID,
                }),
            })
                .then((response) => response.json())
                .then(() => {
                    window.setTimeout(location.href = '/user/invoice', {{index .config "jump_delay"}});
                });
        }
    }).render('#paypal-button-container');

</script>