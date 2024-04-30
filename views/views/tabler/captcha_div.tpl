{{if czeq "turnstile" (index .public_setting "captcha_provider")}}
    <div id="cf-turnstile" class="cf-turnstile" data-sitekey="{{index .captcha "turnstile_sitekey"}}" data-theme="light"></div>
{{end}}
{{if czeq "geetest" (index .public_setting "captcha_provider")}}
    <div id="geetest"></div>
{{end}}
