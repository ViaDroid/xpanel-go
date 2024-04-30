{{template "views/tabler/user/header.tpl" .}}

<div class="page-wrapper">
    <div class="container-xl">
        <div class="page-header d-print-none text-white">
            <div class="row align-items-center">
                <div class="col">
                    <h2 class="page-title">
                        <span class="home-title">工单记录</span>
                    </h2>
                    <div class="page-pretitle my-3">
                        <span class="home-subtitle">你可以在这里查看工单消息并添加回复</span>
                    </div>
                </div>
                {{if ne "closed" .ticket.Status}}
                    <div class="col-auto">
                        <div class="btn-list">
                            <a href="#" class="btn btn-primary" data-bs-toggle="modal"
                               data-bs-target="#add-reply">
                                <i class="icon ti ti-plus"></i>
                                添加回复
                            </a>
                        </div>
                    </div>
                {{end}}
            </div>
        </div>
    </div>
    <div class="page-body">
        <div class="container-xl">
            <div class="row row-cards">
                <div class="col-12">
                    <div class="card">
                        <div class="card-body">
                            <div class="h1 my-2 mb-3">#{{.ticket.Id}} {{.ticket.Title}}</div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="row row-deck my-3">
                <div class="col-4">
                    <div class="card">
                        <div class="card-body">
                            <div class="d-flex align-items-center">
                                <div class="subheader">工单状态</div>
                            </div>
                            <div class="h1 mb-3">{{.ticket.ParseStatus}}</div>
                        </div>
                    </div>
                </div>
                <div class="col-4">
                    <div class="card">
                        <div class="card-body">
                            <div class="d-flex align-items-center">
                                <div class="subheader">工单类型</div>
                            </div>
                            <div class="h1 mb-3">{{.ticket.Type}}</div>
                        </div>
                    </div>
                </div>
                <div class="col-4">
                    <div class="card">
                        <div class="card-body">
                            <div class="card-body">
                                <div class="d-flex align-items-center">
                                    <div class="subheader">工单开启时间</div>
                                </div>
                                <div class="h1 mb-3">{{.ticket.DatetimeStr}}</div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="row justify-content-center my-3">
                <div class="col-12">
                    <div class="card">
                        <div class="card-body">
                            <div class="divide-y">
                                {{range  $comment := .comments}}
                                    <div>
                                        <div class="row">
                                            <div class="col">
                                                <div>
                                                    {{ $comment.comment}}
                                                </div>
                                                <div class="text-secondary my-1">{{$comment.commenter_name}}
                                                    回复于 {{$comment.datetime}}</div>
                                            </div>
                                            <div class="col-auto">
                                                <div>
                                                    # {{$comment.comment_id | inc}}
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                {{end}}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <div class="modal modal-blur fade" id="add-reply" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered modal-dialog-scrollable" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">添加回复</h5>
                    <button class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <div class="mb-3">
                        <textarea id="reply-comment" class="form-control" rows="15" placeholder="请输入回复内容"></textarea>
                    </div>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn me-auto" data-bs-dismiss="modal">取消</button>
                    <button id="reply" class="btn btn-primary" data-bs-dismiss="modal"
                            hx-put="/user/ticket/{{.ticket.Id}}" hx-swap="none"
                            hx-vals='js:{ comment: document.getElementById("reply-comment").value }'>
                        回复
                    </button>
                </div>
            </div>
        </div>
    </div>

    {{template "views/tabler/user/footer.tpl" .}}
