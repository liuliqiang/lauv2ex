lauv2ex
=======

liqiang 的 v2ex 命令行客户端，使用 Go 语言编写，支持简单扩展成各种 App 端和 Web 端。

官方文档
========

`使用手册 <https://liqiang.io/opensource/go/v2ex>`_

`开发手册 <https://liqiang.io/opensource/go/v2ex/dev>`_

功能支持说明
============

社区情况获取
------------------

主题数和用户数获取
^^^^^^^^^^^^^^^^^^^^^
.. code-block:: shell

    /home/yetship/lauv2ex (master ✔) ᐅ v2ex site stat
    Max Topic: 434498
    Max User: 296197

社区描述获取
^^^^^^^^^^^^^^^^^^^^^
.. code-block:: shell

    /home/yetship/lauv2ex (master ✔) ᐅ v2ex site info
    Doamin: www.v2ex.com
    Title: V2EX
    Slogan: way to explore
    Description: 创意工作者们的社区

主题获取
----------------

获取当前热门主题
^^^^^^^^^^^^^^^^

.. code-block:: shell

    /home/yetship/lauv2ex (master ✔) ᐅ v2ex topic hot
    0: 程序员求推荐生产工具(键盘)
    1: 作为一名程序员，从今天开始戒烟。。。
    2: 客厅瓷砖上不上墙？
    3: mweb 的作者你出来
    4: 再过三四十年，那时候的春节会是什么样子。
    5: 小米生活里面的东西为什么可以 5 块钱还包邮？
    6: 微信朋友圈为啥不搞一个分类功能
    7: 北京北边有玩德州扑克的组织嘛
    8: 最近有个项目每晚加到 10 点且周六日都不休

获取当前最新主题
^^^^^^^^^^^^^^^^

.. code-block:: shell

    /home/yetship/lauv2ex (master ✔) ᐅ v2ex topic late
    nb: index  rep title
     0: 434499   0  请问通过网上的教育邮箱撸的谷歌网盘是不是会被收回？
     1: 434497   1  微信群组（北京积极有意义有趣前沿生活方式）
     2: 434495   4  我的身份证在微信被人盗用了吗
     3: 434494   1  为什么 TelegramX 已经禁用了 AutoPlay GIF, 进入频道还是一堆一堆的 gif 在下载和自动播放啊
     4: 434493  10  Flutter 官网上的这张图
     5: 434491   0  有没有代码扫描工具能扫描 txt 文档的？
     6: 434485   5  Chrome 拖出一个 Tab 变成新窗口的动作，是开了一个新窗口 还是 一个新进程？
     7: 434484   6  Mac terimal bash 默认走代理
     8: 434483   1  iPhone6 不小心点了优化照片存储，然后重新设置为保存原件在本机，发现有些照片还是没办法下载下来。。
     9: 434482   2  转转进来
    10: 434481   7  有没有一种爬虫服务，只需要我指定网站和规则，就可以定时爬数据，并且可以提供我指定格式的 json api 的
    11: 434480  16  如何看待公司对于业绩不好的部门没有参加年会资格这件事？

查看某一个主题
^^^^^^^^^^^^^^

.. code-block:: shell

    /home/yetship/lauv2ex (master ✔) ᐅ v2ex topic 434499
    Title: 请问通过网上的教育邮箱撸的谷歌网盘是不是会被收回？
    Url: http://www.v2ex.com/t/434499
    Author: anacn
    Content:
    Replics: 0

查看某一主题的回复
^^^^^^^^^^^^^^^^^^

.. code-block:: shell

    /home/yetship/lauv2ex (master ✔) ᐅ v2ex topic rep 434480
    Id: 0,  User: hsuan,  Tks: 0
    Content: 要我的话求之不得，年会有什么好参加的
    -------------
    Id: 1,  User: zhaolion,  Tks: 1
    Content: @hsuan 重点不在于可以不参加，而是被区别对待了。邀请你，你不去和不让你去，这是有区别的
    -------------
    Id: 2,  User: luoyou1014,  Tks: 0
    Content: 区别对待是正常的，不让参加年会是明面上的区别对待，年终奖少是暗地里的区别对待。

    如果年终奖不被区别对待，那我宁愿不参加年会。
    -------------

版权申明
========

Copyright (C) 2018  Liqiang Liu

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.