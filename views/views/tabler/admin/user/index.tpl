{{template "views/tabler/admin/header.tpl" .}}

<div class="page-wrapper">
    <div class="container-xl">
        <div class="page-header d-print-none text-white">
            <div class="row align-items-center">
                <div class="col">
                    <h2 class="page-title">
                        <span class="home-title">用户列表</span>
                    </h2>
                    <div class="page-pretitle my-3">
                        <span class="home-subtitle">
                            系统中所有用户的列表
                        </span>
                    </div>
                </div>
                <div class="col-auto">
                    <div class="btn-list">
                        <button href="#" class="btn btn-primary" data-bs-toggle="modal"
                                data-bs-target="#create-dialog">
                            <i class="icon ti ti-plus"></i>
                            创建
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

    <div class="modal modal-blur fade" id="create-dialog" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered modal-dialog-scrollable" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">添加用户</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    {{range  $from := (index .details "create_dialog")}}
                        {{if czeq "input" (index $from "Type")}}
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">{{index $from "Info"}}</label>
                                <div class="col">
                                    <input id="{{index $from "Id"}}" type="text" class="form-control"
                                           placeholder="{{index $from "Placeholder"}}">
                                </div>
                            </div>
                        {{end}}
                        {{if czeq "textarea" (index $from "Type")}}
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">{{index $from "Info"}}</label>
                                <textarea id="{{index $from "Id"}}" class="col form-control" rows="{{index $from "rows"}}"
                                          placeholder="{{index $from "Placeholder"}}"></textarea>
                            </div>
                        {{end}}
                        {{if czeq "select" (index $from "Type")}}
                            <div class="form-group mb-3 row">
                                <label class="form-label col-3 col-form-label">{{index $from "Info"}}</label>
                                <div class="col">
                                    <select id="{{index $from "Id"}}" class="col form-select">
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
                    <button id="create-button" type="button" class="btn btn-primary" data-bs-dismiss="modal">添加
                    </button>
                </div>
            </div>
        </div>
    </div>

    <script>
        let table = new DataTable('#data-table', {
            ajax: {
                url: '/admin/user/ajax',
                type: 'POST',
                dataSrc: 'users'
            },
            "autoWidth": false,
            'iDisplayLength': 10,
            'scrollX': true,
            'order': [
                [1, 'asc']
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
                    targets: [0, 6, 7],
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

        function loadTable() {
            table;
        }

        $("#create-button").click(function () {
            $.ajax({
                type: "POST",
                url: "/admin/user/create",
                dataType: "json",
                data: {
                    {{range  $from := (index .details "create_dialog")}}
                        {{index $from "Id"}}: $('#{{index $from "Id"}}').val(),
                    {{end}}
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
        });

        function deleteUser(user_id) {
            $('#notice-message').text('确定删除此用户？');
            $('#notice-dialog').modal('show');
            $('#notice-confirm').off('click').on('click', function () {
                $.ajax({
                    url: "/admin/user/" + user_id,
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

        function reloadTableAjax() {
            table.ajax.reload(null, false);
        }

        loadTable();
    </script>

{{template "views/tabler/admin/footer.tpl" .}}
