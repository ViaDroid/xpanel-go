{{template "views/tabler/admin/header.tpl" .}}

<div class="page-wrapper">
    <div class="container-xl">
        <div class="page-header d-print-none text-white">
            <div class="row align-items-center">
                <div class="col">
                    <h2 class="page-title">
                        <span class="home-title">审计规则</span>
                    </h2>
                    <div class="page-pretitle my-3">
                        <span class="home-subtitle">查看站点中的审计规则</span>
                    </div>
                </div>
                <div class="col-auto">
                    <div class="btn-list">
                        <button href="#" class="btn btn-primary" data-bs-toggle="modal"
                                data-bs-target="#add-detect-dialog">
                            <i class="icon ti ti-plus"></i>
                            添加审计规则
                        </button>
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

    <div class="modal modal-blur fade" id="add-detect-dialog" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered modal-dialog-scrollable" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">添加审计规则</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    {{range  $from := (index .details "add_dialog")}}
                        {{if czeq "input" (index $from "type")}}
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">{{index $from "info"}}</label>
                                <div class="col">
                                    <input id="{{index $from "id"}}" type="text" class="form-control"
                                           placeholder="{{index $from "placeholder"}}">
                                </div>
                            </div>
                        {{end}}
                        {{if czeq "textarea" (index $from "type")}}
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">{{index $from "info"}}</label>
                                <textarea id="{{index $from "id"}}" class="col form-control" rows="{{index $from "rows"}}"
                                          placeholder="{{index $from "placeholder"}}"></textarea>
                            </div>
                        {{end}}
                        {{if czeq "select" (index $from "type")}}
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">{{index $from "info"}}</label>
                                <div class="col">
                                    <select id="{{index $from "id"}}" class="col form-select">
                                        {{range $key, $value := (index $from "select")}}
                                            <option value="{{$key}}">{{$value}}</option>
                                        {{end}}
                                    </select>
                                </div>
                            </div>
                        {{end}}
                    {{end}}
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn me-auto" data-bs-dismiss="modal">取消</button>
                    <button id="add-detect-button" type="button" class="btn btn-primary" data-bs-dismiss="modal">提交
                    </button>
                </div>
            </div>
        </div>
    </div>

    <script>
        let table = new DataTable('#data-table', {
            ajax: {
                url: '/admin/detect/ajax',
                type: 'POST',
                dataSrc: 'rules'
            },
            "autoWidth": false,
            'iDisplayLength': 10,
            'scrollX': true,
            'order': [
                [0, 'desc']
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
                },
            ],
            "dom": "<'row px-3 py-3'<'col-sm-12 col-md-6'l><'col-sm-12 col-md-6'f>>" +
                "<'row'<'col-sm-12'tr>>" +
                "<'row card-footer d-flex align-items-center'<'col-sm-12 col-md-5'i><'col-sm-12 col-md-7'p>>",
            language: {
                "sProcessing": "处理中...",
                "sLengthMenu": "显示 _MENU_ 条",
                "sZeroRecords": "没有匹配结果",
                "sInfo": "第 _START_ 至 _END_ 项结果，共 _TOTAL_ 项",
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
                    "sPrevious": "<i class=\"ti ti-arrow-left\"></i>",
                    "sNext": "<i class=\"ti ti-arrow-right\"></i>",
                    "sLast": "末页"
                },
                "oAria": {
                    "sSortAscending": ": 以升序排列此列",
                    "sSortDescending": ": 以降序排列此列"
                }
            }
        });

        $("#add-detect-button").click(function () {
            $.ajax({
                type: "POST",
                url: "/admin/detect/add",
                dataType: "json",
                data: {
                    {{range  $from := (index .details "add_dialog")}}
                    {{index $from "id"}}: $('#{{index $from "id"}}').val(),
                    {{end}}
                },
                success: function (data) {
                    if (data.ret == 1) {
                        $('#success-message').text(data.msg);
                        $('#success-dialog').modal('show');
                        reloadTableAjax();
                    } else {
                        $('#fail-message').text(data.msg);
                        $('#fail-dialog').modal('show');
                    }
                }
            })
        });

        function deleteRule(rule_id) {
            $('#notice-message').text('确定删除此审计规则？');
            $('#notice-dialog').modal('show');
            $('#notice-confirm').off('click').on('click', function () {
                $.ajax({
                    url: "/admin/detect/" + rule_id,
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

        function loadTable() {
            table;
        }

        function reloadTableAjax() {
            table.ajax.reload(null, false);
        }

        loadTable();
    </script>

    {{template "views/tabler/admin/footer.tpl" .}}
