-- ----------------------------
-- 背景图片ID
-- ----------------------------
SET @instance_background_img = 13;

-- ----------------------------
-- screen_instance（大屏实例）
-- ----------------------------
SET @version = 1;

INSERT INTO `screen_instance` (
    `add_time`,
    `add_user`,
    `edit_time`,
    `edit_user`,
    `del_flag`,
    `instance_title`,
    `instance_background_color`,
    `instance_background_img`,
    `instance_height`,
    `instance_width`,
    `instance_version`,
    `instance_view_img`
)
VALUES
(
    '2020-02-28 13:27:30',
    '1',
    '1',
    '2020-02-28 13:27:30',
    '1',
    '#263546',
    @instance_background_img,
    '1080',
    '零售行业数据可视化',
    '',
    '1920',
    @version
);

-- ----------------------------
-- \"database\":9
-- 更换为自己的ID
-- ----------------------------

SET @instance_id = @@IDENTITY;

-- ----------------------------
-- 顶部大标题
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '50',
    'chart0',
    '8000',
    '{\"title\":\"零售行业数据可视化\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":43,\"fontWeight\":\"normal\",\"textAlign\":\"center\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '480',
    '720',
    '25',
    @version
);

-- ----------------------------
-- 销售额数据
-- ----------------------------

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"\",\"database\":\"\",\"fileName\":\"\",\"sql\":\"\",\"x\":\"\",\"name\":\"\"}',
    'titleText',
    'false',
    '[]',
    '25',
    'chart1',
    '8000',
    '{\"title\":\"销售额数据\",\"fontColor\":\"rgba(255,255,255,1)\",\"fontSize\":17,\"fontWeight\":\"normal\",\"textAlign\":\"center\",\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '585',
    '30',
    '122',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select * from template_5_sd_histogram\",\"y\":\"y\",\"x\":\"x\",\"chartType\":\"histogramGradient\"}',
    'histogramGradient',
    'false',
    '{}',
    '240',
    'chart2',
    '8000',
    '{\"title\":{\"text\":\"渐变柱状图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"legend\":{\"show\":false},\"backgroundColor\":\"\",\"grid\":{\"left\":\"0\",\"right\":\"5%\",\"bottom\":\"0\",\"containLabel\":true,\"top\":\"5%\"},\"xAxis\":{\"data\":[\"苹果\",\"三星\",\"华为\",\"oppo\",\"vivo\",\"小米\"],\"name\":\"\",\"axisTick\":{\"show\":false},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLabel\":{\"show\":true,\"inside\":false,\"textStyle\":{\"color\":\"#fff\"},\"rotate\":0,\"interval\":0},\"z\":10},\"yAxis\":{\"name\":\"\",\"axisTick\":{\"show\":false},\"axisLine\":{\"show\":true,\"lineStyle\":{\"color\":\"#ffffff\"}},\"splitLine\":{\"show\":false,\"lineStyle\":{\"color\":\"#ffffff\"}},\"axisLabel\":{\"show\":true,\"textStyle\":{\"color\":\"#fff\"}}},\"dataZoom\":[{\"type\":\"inside\"}],\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"series\":[{\"type\":\"bar\",\"itemStyle\":{\"normal\":{\"color\":\"rgba(255,255,255,0)\"}},\"barWidth\":\"30%\",\"barGap\":\"-100%\",\"barCategoryGap\":\"0\",\"data\":[[7739],[7739],[7739],[7739],[7739],[7739]],\"animation\":false},{\"type\":\"bar\",\"barWidth\":\"40%\",\"itemStyle\":{\"normal\":{\"color\":{\"x\":0,\"y\":0,\"x2\":0,\"y2\":1,\"type\":\"linear\",\"global\":false,\"colorStops\":[{\"offset\":0,\"color\":\"#83bff6\"},{\"offset\":0.5,\"color\":\"#188df0\"},{\"offset\":1,\"color\":\"#188df0\"}]}},\"emphasis\":{\"color\":{\"x\":0,\"y\":0,\"x2\":0,\"y2\":1,\"type\":\"linear\",\"global\":false,\"colorStops\":[{\"offset\":0,\"color\":\"#2378f7\"},{\"offset\":0.7,\"color\":\"#2378f7\"},{\"offset\":1,\"color\":\"#83bff6\"}]}}},\"data\":[6532,7739,6160,7373,7755,6420]}]}',
    'false',
    '330',
    '35',
    '180',
    @version
);

INSERT INTO `chart_item` (
    `instance_id`,
    `item_chart_data`,
    `item_chart_type`,
    `item_choose`,
    `item_data`,
    `item_height`,
    `item_i`,
    `item_interval`,
    `item_option`,
    `item_refresh`,
    `item_width`,
    `item_x`,
    `item_y`,
    `item_version`
)
VALUES
(
    @instance_id,
    '{\"dataSourceType\":\"DataBase\",\"database\":9,\"sql\":\"select name,value from template_5_sd_pie\",\"name\":\"name\",\"value\":\"value\",\"chartType\":\"pie2D\"}',
    'pie2D',
    'false',
    '[]',
    '260',
    'chart3',
    '8000',
    '{\"lunbo\":\"0\",\"title\":{\"text\":\"2D饼图\",\"left\":\"center\",\"textStyle\":{\"color\":\"#fff\"},\"show\":false},\"tooltip\":{\"trigger\":\"item\",\"formatter\":\"\"},\"legend\":{\"show\":false,\"orient\":\"vertical\",\"top\":0,\"left\":0,\"data\":[\"直接访问\",\"邮件营销\",\"联盟广告\",\"视频广告\",\"搜索引擎\"],\"textStyle\":{\"color\":\"#fff\"}},\"color\":[\"#24CBFF\",\"#7760F6\",\"#484CDC\",\"#711D8E\",\"#2D3C98\"],\"series\":[{\"name\":\"访问来源\",\"type\":\"pie\",\"radius\":\"50%\",\"center\":[\"50%\",\"50%\"],\"data\":[{\"name\":\"直接访问\",\"value\":\"360\"},{\"name\":\"邮件营销\",\"value\":\"335\"},{\"name\":\"联盟广告\",\"value\":\"310\"},{\"name\":\"视频广告\",\"value\":\"275\"},{\"name\":\"搜索引擎\",\"value\":\"235\"}],\"roseType\":\"angle\",\"label\":{\"normal\":{\"show\":true,\"position\":\"outside\",\"textStyle\":{\"color\":\"#ffffff\"}}},\"labelLine\":{\"normal\":{\"show\":true,\"lineStyle\":{\"color\":\"#ffffff\"},\"smooth\":0.2,\"length\":10,\"length2\":20}},\"itemStyle\":{\"normal\":{\"color\":\"\",\"shadowBlur\":200,\"shadowColor\":\"rgba(0,0,0,0)\"}},\"animationType\":\"scale\",\"animationEasing\":\"elasticOut\",\"avoidLabelOverlap\":true}],\"toolbox\":{\"feature\":{\"saveAsImage\":{}}},\"backgroundColor\":\"rgba(255,255,255,0)\"}',
    'false',
    '300',
    '330',
    '180',
    @version
);
