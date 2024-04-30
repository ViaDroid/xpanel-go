{{if czeq "turnstile" (index .public_setting "captcha_provider")}}
    <script src="https://challenges.cloudflare.com/turnstile/v0/api.js" async defer></script>
{{end}}
{{if czeq "geetest" (index .public_setting "captcha_provider")}}
    <script src="https://static.geetest.com/v4/gt4.js"></script>
    <script>
        let geetest_result = '';
        initGeetest4({
            captchaId: '{{.geetest_id}}',
            product: 'float',
            language: "zho",
            riskType: 'slide'
        }, function (geetest) {
            geetest.appendTo("#geetest");
            geetest.onSuccess(function () {
                geetest_result = geetest.getValidate();
            });
        });
    </script>
{{end}}
