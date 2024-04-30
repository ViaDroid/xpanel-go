{{template "views/tabler/admin/header.tpl" .}}

<link rel="stylesheet" href="//{{index .config "jsdelivr_url"}}/npm/flatpickr/dist/flatpickr.min.css">
{{if .user.IsDarkMode}}
    <link rel="stylesheet" href="//{{index .config "jsdelivr_url"}}/npm/flatpickr/dist/themes/dark.min.css">
{{end}}
<script src="//{{index .config "jsdelivr_url"}}/npm/flatpickr"></script>
<script src="//{{index .config "jsdelivr_url"}}/npm/flatpickr/dist/l10n/zh.js"></script>

<div class="page-wrapper">
    <div class="container-xl">
        <div class="page-header d-print-none text-white">
            <div class="row align-items-center">
                <div class="col">
                    <h2 class="page-title">
                        <span class="home-title">优惠码</span>
                    </h2>
                    <div class="page-pretitle my-3">
                        <span class="home-subtitle">
                            查看并管理优惠码
                        </span>
                    </div>
                </div>
                <div class="col-auto">
                    <div class="btn-list">
                        <a href="#" class="btn btn-primary" data-bs-toggle="modal"
                           data-bs-target="#create-dialog">
                            <i class="icon ti ti-plus"></i>
                            创建
                        </a>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="page-body">
        <div class="container-xl">
            <div class="row row-deck row-cards">
                <div class="col-12">
                    <div class="card">
                        <div class="table-responsive">
                            <table id="data-table" class="table card-table table-vcenter text-nowrap datatable">
                                <thead>
                                <tr>
                                    {{range $key, $value := (index .details "field")}}
                                        <th>{{$value.Value}}</th>
                                    {{end}}
                                </tr>
                                </thead>
                            </table>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <div class="modal modal-blur fade" id="create-dialog" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered modal-dialog-scrollable" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">优惠码内容</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    {{range  $detail := (index .details "create_dialog")}}
                        {{if czeq "input" (index $detail "Type")}}
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">{{index $detail "Info"}}</label>
                                <div class="col">
                                    <input id="{{index $detail "Id"}}" type="text" class="form-control"
                                           placeholder="{{index $detail "Placeholder"}}">
                                </div>
                            </div>
                        {{end}}
                        {{if czeq "textarea" (index $detail "Type")}}
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">{{index .detail "info"}}</label>
                                <textarea id="{{index $detail "id"}}" class="col form-control" rows="{{index $detail "rows"}}"
                                          placeholder="{{index $detail "Placeholder"}}"></textarea>
                            </div>
                        {{end}}
                        {{if czeq "select" (index $detail "Type")}}
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">{{index $detail "Info"}}</label>
                                <div class="col">
                                    <select id="{{index $detail "Id"}}" class="col form-select">
                                        {{range $key, $value := (index $detail "select")}}
                                            <option value="{{$key}}">{{$value}}</option>
                                        {{end}}
                                    </select>
                                </div>
                            </div>
                        {{end}}
                    {{end}}
                    <div class="form-group mb-3 row">
                        <label class="form-label col-3 col-form-label">过期时间（留空则为不限制）</label>
                        <div class="col">
                            <input id="expire_time" type="text" class="form-control"
                                   placeholder="">
                        </div>
                    </div>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn me-auto" data-bs-dismiss="modal">取消</button>
                    <button id="create-button" onclick="createCoupon()"
                            type="button" class="btn btn-primary" data-bs-dismiss="modal">创建
                    </button>
                </div>
            </div>
        </div>
    </div>

    <script>
        flatpickr("#expire_time", {
            enableTime: true,
            //dateFormat: "U",
            time_24hr: true,
            minDate: "today",
            locale: "zh"
        });

        let table = new DataTable('#data-table', {
            ajax: {
                url: '/admin/coupon/ajax',
                type: 'POST',
                dataSrc: 'coupons'
            },
            "autoWidth": false,
            'iDisplayLength': 10,
            'scrollX': true,
            'order': [
                [1, 'desc']
            ],
            columns: [
                {{range $key, $value := (index .details "field")}}
                {
                    data: '{{$value.Key}}'
                },
                {{end}}
            ],
            "columnDefs": [
                {
                    targets: [0],
                    orderable: false
                }
            ],
            "dom": "<'row px-3 py-3'<'col-sm-12 col-md-6'l><'col-sm-12 col-md-6'f>>" +
                "<'row'<'col-sm-12'tr>>" +
                "<'row card-footer d-flex d-flexalign-items-center'<'col-sm-12 col-md-5'i><'col-sm-12 col-md-7'p>>",
            language: {
                "sProcessing": "处理中...",
                "sLengthMenu": "显示 _MENU_ 条",
                "sZeroRecords": "没有匹配结果",
                "sInfo": "第 _START_ 至 _END_ 项结果，共 _TOTAL_项",
                "sInfoEmpty": "第 0 至 0 项结果，共 0 项",
                "sInfoFiltered": "(在 _MAX_ 项中查找)",
                "sInfoPostFix": "",
                "sSearch": "<i class=\"ti ti-search\"></i> ",
                "sUrl": "",
                "sEmptyTable": "表中数据为空",
                "sLoadingRecords": "载入中...",
                "sInfoThousands": ",",
                "oPaginate": {
                    "sFirst": "首页",
                    "sPrevious": "<i class=\"titi-arrow-left\"></i>",
                    "sNext": "<i class=\"ti ti-arrow-right\"><i>",
                    "sLast": "末页"
                },
                "oAria": {
                    "sSortAscending": ": 以升序排列此列",
                    "sSortDescending": ": 以降序排列此列"
                }
            },
        });

        function loadTable() {
            table;
        }

        function createCoupon() {
            $.ajax({
                url: '/admin/coupon',
                type: 'POST',
                dataType: "json",
                data: {
                    {{range  $detail := (index .details "create_dialog")}}
                    {{index $detail "Id"}}: $('#{{index $detail "Id"}}').val(),
                    {{end}}
                    expire_time: $('#expire_time').val(),
                },
                success: function (data) {
                    if (data.ret == 1) {
                        $('#success-noreload-message').text(data.msg);
                        $('#success-noreload-dialog').modal('show');
                        reloadTableAjax();
                    } else {
                        $('#fail-message').text(data.msg);
                        $('#fail-dialog').modal('show');
                    }
                }
            })
        }

        function deleteCoupon(coupon_id) {
            $('#notice-message').text('确定删除此优惠码？');
            $('#notice-dialog').modal('show');
            $('#notice-confirm').off('click').on('click', function () {
                $.ajax({
                    url: "/admin/coupon/" + coupon_id,
                    type: 'DELETE',
                    dataType: "json",
                    success: function (data) {
                        if (data.ret == 1) {
                            $('#success-noreload-message').text(data.msg);
                            $('#success-noreload-dialog').modal('show');
                            reloadTableAjax();
                        } else {
                            $('#fail-message').text(data.msg);
                            $('#fail-dialog').modal('show');
                        }
                    }
                })
            });
        }

        function disableCoupon(coupon_id) {
            $('#notice-message').text('确定禁用此优惠码？');
            $('#notice-dialog').modal('show');
            $('#notice-confirm').off('click').on('click', function () {
                $.ajax({
                    url: "/admin/coupon/" + coupon_id + "/disable",
                    type: 'POST',
                    dataType: "json",
                    success: function (data) {
                        if (data.ret == 1) {
                            $('#success-noreload-dialog').text(data.msg);
                            $('#success-noreload-message').modal('show');
                            reloadTableAjax();
                        } else {
                            $('#fail-message').text(data.msg);
                            $('#fail-dialog').modal('show');
                        }
                    }
                })
            });
        }

        function reloadTableAjax() {
            table.ajax.reload(null, false);
        }

        loadTable();
    </script>

    {{template "views/tabler/admin/footer.tpl" .}}
