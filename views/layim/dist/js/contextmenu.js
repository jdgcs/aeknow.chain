layui.define(['layer', 'element'], function (exports) {

    const $ = layui.jquery
        , layer = layui.layer
        , element = layui.element
        , device = layui.device()
        , stope = layui.stope;


    let defaults = [{
        menu: [{
            text: "菜单一",
            callback: function (t) {
            }
        }, {
            text: "菜单二",
            callback: function (t) {
            }
        }],
        target: $('.layim-list-friend')
    }];

    const menu = {
        init: function (options) {
            defaults = $.extend(defaults, options);
            layui.each(defaults, function (index, item){
                menu.menuClick(item);
            })
        },
        hide: function () {
            layer.closeAll('tips');
        },
        menuClick: function (options) {
            let target = options.target;

            $(target).on('contextmenu', function (event) {
                if (event.target != this) return false;

                let lis = '';
                layui.each(options.menu, function (index, item) {
                    lis += '<li class="ui-context-menu-item"><a href="javascript:void(0);" ><span>' + item.text + '</span></a></li>';
                });

                let html = '<ul id="contextmenu">' + lis + '</ul>';
                layer.tips(html, target, {
                    tips: 1,
                    time: 0,
                    anim: 5,
                    fixed: true,
                    skin: "layui-box layui-layim-contextmenu",
                    success: function (layero, index) {
                        menu.menuChildrenClick(options);
                        const stopEvent = function (event) {
                            stope(event);
                        };
                        layero.off('mousedown', stopEvent).on('mousedown', stopEvent);
                    }
                });
            });

            $(document).off('mousedown', menu.hide).on('mousedown', menu.hide);
        },
        menuChildrenClick: function (options) {
            $(document).off('click').on("click", ".ui-context-menu-item", function () {
                let i = $(this).index();
                layer.closeAll('tips');
                options.menu[i].callback && "function" == typeof options.menu[i].callback && options.menu[i].callback($(this));
                stope(options.menu[i].callback);
            });
        }
    };

    exports('contextmenu', menu);
});
