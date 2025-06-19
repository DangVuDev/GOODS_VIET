import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Form, Input, Button, Alert, Card } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';

const Login: React.FC = () => {
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const onFinish = (values: { username: string; password: string }) => {
    if (values.username === 'ADMIN' && values.password === 'ADMIN') {
      localStorage.setItem('isAuthenticated', 'true');
      navigate('/dashboard');
    } else {
      setError('Tài khoản hoặc mật khẩu không đúng');
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gradient-to-br from-blue-100 to-purple-100 px-4">
      <Card className="w-full max-w-md p-6 card-hover">
        <h2 className="text-2xl font-bold mb-6 text-center text-blue-800">Đăng nhập</h2>
        {error && <Alert message={error} type="error" className="mb-4" />}
        <Form onFinish={onFinish} layout="vertical">
          <Form.Item
            name="username"
            rules={[{ required: true, message: 'Vui lòng nhập tài khoản!' }]}
          >
            <Input prefix={<UserOutlined />} placeholder="Tài khoản" size="large" />
          </Form.Item>
          <Form.Item
            name="password"
            rules={[{ required: true, message: 'Vui lòng nhập mật khẩu!' }]}
          >
            <Input.Password prefix={<LockOutlined />} placeholder="Mật khẩu" size="large" />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" block size="large">
              Đăng nhập
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
};

export default Login;