$excel_path = EXTEND_PATH . 'phpexcel/';

require $excel_path . 'PHPExcel.php';

$excel = new \PHPExcel();
$sheet = $excel->getActiveSheet();
$sheet->setTitle('订单报表');


$list = $model->index('addTime', $data);
$data = $list->toArray();
$excelBodyData = [];

$array = [
    ['序号', '订单号', 'VIP账号', '联系人', '联系电话', '预约服务', '所属门店', '服务类型', '预约次数', '订单金额', '所属服务点', '订单状态', '付款时间'],
];

for ($i = 0; $i <= 9; ++$i) {

    switch ($data['data'][$i]['state']) {

        case 1:
            $data['data'][$i]['state'] = '待付款';
            break;
        case 2:
            $data['data'][$i]['state'] = '待分配';
            break;
        case 3:
            $data['data'][$i]['state'] = '待服务';
            break;
        case 4:
            $data['data'][$i]['state'] = '服务中';
            break;
        case 5:
            $data['data'][$i]['state'] = '服务已完成';
            break;
        case 6:
            $data['data'][$i]['state'] = '已退款';
            break;
        case 7:
            $data['data'][$i]['state'] = '已关闭';
            break;
    }

    $excelBodyData[$i]['id'] = $data['data'][$i]['id'];
    $excelBodyData[$i]['order_number'] = $data['data'][$i]['order_number'];
    $excelBodyData[$i]['vip_number'] = $data['data'][$i]['vip_number'];
    $excelBodyData[$i]['contact_people'] = $data['data'][$i]['contact_people'];
    $excelBodyData[$i]['contact_phone'] = $data['data'][$i]['contact_phone'];
    $excelBodyData[$i]['product'] = (json_decode($data['data'][$i]['product'], true))['name'];
    $excelBodyData[$i]['store'] = $data['data'][$i]['store'];

    $excelBodyData[$i]['type'] = ($data['data'][$i]['type'] === 1) ? '上门' : '预约到店';

    $excelBodyData[$i]['num'] = $data['data'][$i]['num'];

    $excelBodyData[$i]['order_amount'] = $data['data'][$i]['order_amount'];
    $excelBodyData[$i]['service'] = $data['data'][$i]['service'];
    $excelBodyData[$i]['state'] = $data['data'][$i]['state'];
    $excelBodyData[$i]['payment_time'] = funcs_toDatetime($data['data'][$i]['payment_time']);

    array_push($array, $excelBodyData[$i]);
}
            
$sheet->fromArray($array);
$sheet->getColumnDimension('B')->setWidth(30);
$sheet->getColumnDimension('F')->setWidth(30);
$sheet->getColumnDimension('M')->setWidth(30);

// 生成的 Excel 文件的名字
$fileName = md5(md5(md5(rand(10, 999) . 'and' . rand(1001, 9999))));

// 输出到客户端
header('Content-Type: application/vnd.ms-excel');
header('Content-Disposition: attachment;filename="' . $fileName . '.xls"');
header('Cache-Control: max-age=0');
header('Cache-Control: max-age=1');
header('Expires: Mon, 26 Jul 1997 05:00:00 GMT'); // Date in the past
header('Last-Modified: ' . gmdate('D, d M Y H:i:s') . ' GMT'); // always modified
header('Cache-Control: cache, must-revalidate'); // HTTP/1.1
header('Pragma: public'); // HTTP/1.0
$objWriter = new \PHPExcel_Writer_Excel5($excel);
$objWriter->save('php://output');

// 响应
$this->responseJson($list, 0);
