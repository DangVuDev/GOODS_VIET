import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import { Card, Timeline, Spin, Form, Input, Button, Tag, Space, Checkbox, Modal, QRCode } from 'antd';
import { ShareAltOutlined } from '@ant-design/icons';
import { getProductById, updateProductStatus } from '../services/api';
import type { Product } from '../types';

const ProductDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const [product, setProduct] = useState<Product | null>(null);
  const [isSold, setIsSold] = useState(false);
  const [form] = Form.useForm();
  const [isShareModalOpen, setIsShareModalOpen] = useState(false);

  useEffect(() => {
    const fetchProduct = async () => {
      try {
        if (id) {
          const data = await getProductById(parseInt(id));
          setProduct(data);
          const latestStatus = data.statuses[data.statuses.length - 1]?.status;
          if (latestStatus === 'Đã bán') {
            setIsSold(true);
          }
        }
      } catch (error) {
        console.error('Error fetching product:', error);
      }
    };
    fetchProduct();
  }, [id]);

  const onFinish = async (values: { status: string; details: string; isSold?: boolean }) => {
    if (!id) return;
    try {
      const statusToSubmit = values.isSold ? 'Đã bán' : values.status;
      await updateProductStatus(parseInt(id), {
        status: statusToSubmit,
        details: values.details,
        timestamp: new Date().toISOString(),
      });
      const updatedProduct = await getProductById(parseInt(id));
      setProduct(updatedProduct);
      if (updatedProduct.statuses[updatedProduct.statuses.length - 1]?.status === 'Đã bán') {
        setIsSold(true);
      }
      form.resetFields();
    } catch (error) {
      console.error('Error updating status:', error);
    }
  };

  const handleCheckboxChange = (e: any) => {
    const checked = e.target.checked;
    form.setFieldsValue({ status: checked ? 'Đã bán' : '' });
  };

  const handleShareClick = () => {
    setIsShareModalOpen(true);
  };

  const handleModalClose = () => {
    setIsShareModalOpen(false);
  };

  const shareUrl = product?.qrCodeUrl || window.location.href;

  if (!product) return (
    <div className="min-h-screen flex justify-center items-center bg-gradient-to-br from-indigo-100 via-purple-100 to-pink-100">
      <Spin size="large" className="text-indigo-600" />
    </div>
  );

  return (
    <div className="min-h-screen flex justify-center items-center bg-gradient-to-br from-indigo-100 via-purple-100 to-pink-100 px-4 py-8">
      <Card
        className="w-full max-w-lg transform transition-all duration-300 hover:scale-105 hover:shadow-2xl bg-white/95 backdrop-blur-sm border-none rounded-2xl overflow-hidden"
        cover={
          <div className="relative w-full aspect-video overflow-hidden">
            <img
              src={product.imageURL || '/placeholder.png'}
              alt="Product"
              className="w-full h-full object-cover transition-transform duration-500 hover:scale-110"
            />
            <div className="absolute top-4 left-4 bg-gradient-to-r from-indigo-600 to-purple-600 text-white px-4 py-2 rounded-full text-sm font-semibold shadow-md">
              Mã sản phẩm: {product.productId}
            </div>
            <Button
              icon={<ShareAltOutlined />}
              className="absolute top-4 right-4 bg-white/90 text-indigo-600 hover:text-purple-600 border-none rounded-full w-10 h-10 flex items-center justify-center shadow-md transition-all duration-200 hover:bg-indigo-100"
              onClick={handleShareClick}
            />
          </div>
        }
      >
        <div className="p-6">
          <h2 className="text-2xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-indigo-600 to-purple-600 mb-4 text-center">
            {product.name}
          </h2>
          <p className="text-base text-gray-600 mb-6 text-center">
            <strong>Nhà sản xuất:</strong> {product.manufacturer}
          </p>
          <h3 className="text-lg font-semibold text-indigo-600 mb-4 text-center">
            Lịch sử trạng thái
          </h3>
          <Timeline className="px-4">
            {product.statuses.map((status, index) => (
              <Timeline.Item
                key={index}
                color="indigo"
                className="transition-all duration-300 hover:bg-indigo-50 rounded-lg p-2"
              >
                <Space direction="vertical" size="small">
                  <Tag
                    color="blue"
                    className="text-sm font-medium bg-blue-100 text-blue-700 border-none"
                  >
                    {status.status}
                  </Tag>
                  <p className="text-gray-600 text-sm">{status.details}</p>
                  <p className="text-gray-400 text-xs">
                    {new Date(status.timestamp).toLocaleString()}
                  </p>
                </Space>
              </Timeline.Item>
            ))}
          </Timeline>

          {localStorage.getItem('isAuthenticated') === 'true' && !isSold && (
            <div className="mt-8">
              <h3 className="text-lg font-semibold text-indigo-600 mb-4 text-center">
                Cập nhật trạng thái
              </h3>
              <Form
                form={form}
                onFinish={onFinish}
                layout="vertical"
                className="space-y-4"
              >
                <Form.Item
                  label={<span className="text-indigo-600 font-medium">Trạng thái</span>}
                  name="status"
                  rules={[{ required: !form.getFieldValue('isSold'), message: 'Vui lòng nhập trạng thái!' }]}
                >
                  <Input
                    className="rounded-lg border-indigo-300 focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 transition-all duration-200"
                    disabled={form.getFieldValue('isSold')}
                  />
                </Form.Item>
                <Form.Item
                  label={<span className="text-indigo-600 font-medium">Chi tiết</span>}
                  name="details"
                  rules={[{ required: true, message: 'Vui lòng nhập chi tiết!' }]}
                >
                  <Input
                    className="rounded-lg border-indigo-300 focus:border-indigo-500 focus:ring focus:ring-indigo-200 focus:ring-opacity-50 transition-all duration-200"
                  />
                </Form.Item>
                <Form.Item name="isSold" valuePropName="checked">
                  <Checkbox
                    onChange={handleCheckboxChange}
                    className="text-indigo-600"
                  >
                    Đã mua
                  </Checkbox>
                </Form.Item>
                <Form.Item>
                  <Button
                    type="primary"
                    htmlType="submit"
                    className="w-full h-10 bg-gradient-to-r from-indigo-600 to-purple-600 hover:from-indigo-700 hover:to-purple-700 text-white font-semibold rounded-lg transition-all duration-200"
                  >
                    Cập nhật trạng thái
                  </Button>
                </Form.Item>
              </Form>
            </div>
          )}

          {isSold && (
            <div className="mt-6 text-center">
              <p className="text-red-500 font-semibold">
                Sản phẩm đã được mua, không thể cập nhật trạng thái.
              </p>
            </div>
          )}
        </div>
      </Card>

      <Modal
        title={<span className="text-lg font-semibold text-indigo-600">Chia sẻ sản phẩm</span>}
        open={isShareModalOpen}
        onCancel={handleModalClose}
        footer={null}
        centered
        className="rounded-2xl"
      >
        <div className="flex flex-col items-center space-y-4 p-4">
          <QRCode
            value={shareUrl}
            size={200}
            className="border border-indigo-200 p-2 rounded-lg"
          />
          <Input
            value={shareUrl}
            readOnly
            className="rounded-lg border-indigo-300 text-center"
          />
          <Button
            type="primary"
            className="bg-gradient-to-r from-indigo-600 to-purple-600 hover:from-indigo-700 hover:to-purple-700 text-white font-semibold rounded-lg"
            onClick={() => navigator.clipboard.writeText(shareUrl)}
          >
            Sao chép liên kết
          </Button>
        </div>
      </Modal>
    </div>
  );
};

export default ProductDetail;