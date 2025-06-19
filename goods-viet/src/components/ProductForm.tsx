import { useState } from 'react';
import { Form, Input, Button, Card, Upload } from 'antd';
import { UploadOutlined } from '@ant-design/icons';
import { addProduct } from '../services/api';
import QRCodeDisplay from './QRCodeDisplay';

const ProductForm: React.FC = () => {
  const [qrCodeData, setQrCodeData] = useState<string | null>(null);
  const [file, setFile] = useState<File | null>(null);
  const [form] = Form.useForm();

  const onFinish = async (values: any) => {
    try {
      const response = await addProduct(values, file);
      setQrCodeData(response.qrCodeUrl);
      form.resetFields();
      setFile(null);
    } catch (error) {
      console.error('Error adding product:', error);
    }
  };

  const handleFileChange = (info: any) => {
    if (info.fileList.length > 0) {
      setFile(info.fileList[0].originFileObj);
    } else {
      setFile(null);
    }
  };

  return (
    <Card className="mb-6 card-hover" title={<h2 className="text-xl font-bold">Thêm sản phẩm mới</h2>}>
      <Form form={form} onFinish={onFinish} layout="vertical" className="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <Form.Item
          label="Tên sản phẩm"
          name="name"
          rules={[{ required: true, message: 'Vui lòng nhập tên sản phẩm!' }]}
        >
          <Input size="large" />
        </Form.Item>
        <Form.Item
          label="Nhà sản xuất"
          name="manufacturer"
          rules={[{ required: true, message: 'Vui lòng nhập nhà sản xuất!' }]}
        >
          <Input size="large" />
        </Form.Item>
        <Form.Item
          label="Trạng thái ban đầu"
          name="initialStatus"
          rules={[{ required: true, message: 'Vui lòng nhập trạng thái!' }]}
        >
          <Input size="large" />
        </Form.Item>
        <Form.Item
          label="Chi tiết ban đầu"
          name="initialDetails"
          rules={[{ required: true, message: 'Vui lòng nhập chi tiết!' }]}
        >
          <Input size="large" />
        </Form.Item>
        <Form.Item label="Hình ảnh" className="sm:col-span-2">
          <Upload
            onChange={handleFileChange}
            beforeUpload={() => false} // Ngăn upload tự động
            accept="image/*"
            maxCount={1}
          >
            <Button icon={<UploadOutlined />}>Tải lên hình ảnh</Button>
          </Upload>
        </Form.Item>
        <Form.Item className="sm:col-span-2">
          <Button type="primary" htmlType="submit" size="large" block>
            Thêm sản phẩm
          </Button>
        </Form.Item>
      </Form>
      {qrCodeData && <QRCodeDisplay qrCodeUrl={qrCodeData} />}
    </Card>
  );
};

export default ProductForm;