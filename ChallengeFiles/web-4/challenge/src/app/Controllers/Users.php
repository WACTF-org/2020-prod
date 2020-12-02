<?php namespace App\Controllers;

use App\Models\UsersModel;
use CodeIgniter\Controller;

class Users extends BaseController
{
    public function index()
    {
        if($this->request->getCookie('securelyLoggedInUser') == 'superadmin'){
            # They are already logged in and don't need to be here.
            return redirect()->route('/');
        }else{
           # They need an account to continue
           return redirect()->to(base_url('users/signup_required'));
        }
    }

    public function login(){
        # change this to be my new special cookie later on.
        if($this->request->getCookie('securelyLoggedInUser') == 'superadmin'){
            # They are already logged in and don't need to be here.
            return redirect()->route('/');
        }else{
            if($this->request->getMethod() === 'post'){
                $model = new UsersModel();
                $username = $this->request->getPost('username');
                $password = $model->getUserPass($username);
                return $password;
            }else{
                # Display the login page
                $data['title'] = 'Login';
                echo view('templates/header', $data);
                echo view('users/login');
                echo view('templates/footer');
            }
        }
    }

    public function register(){
        if($this->request->getCookie('securelyLoggedInUser') == 'superadmin'){
            # They are already logged in and don't need to be here.
            return redirect()->route('/');
        }else{
            $data['title'] = 'Register';
            echo view('templates/header', $data);
            if($this->request->getMethod() === 'post'){
                # Tell them to get lost
                if($this->validate(['username' => 'required|min_length[1]',
                                    'password' => 'required|min_length[8]'])){
                    # If their request was well formed.
                    $password = password_hash($this->request->getPost('password'), PASSWORD_DEFAULT);
                    $username = $this->request->getPost('username');
                    $model = new UsersModel();
                    # Check if submitted credentials match a stored user
                    if($model->getUserId($username) !== NULL){
                        echo view('users/register', ['errors' => 'Username is in use.']);
                    }else{
                        echo view('users/register', ['errors' => 'User account creation not currently supported! Please contact your administrator if you need an account!']);
                    }
                }else{
                    # They submitted data that doesn't pass validation rules.
                    echo view('users/register', ['errors' => $this->validator->getErrors()]);
                }
            }else{
                # Display the register page
                echo view('users/register');
            }
            echo view('templates/footer');
        }
    }

    public function logout(){
        $this->session->remove(['username', 'id', 'is_logged_in']);
        $this->session->destroy(); 
        return redirect('users/login');
    }

    public function signup_required(){
        $data['title'] = 'Sign Up Required';
        echo view('templates/header', $data);
        echo view('users/signup_required');
        echo view('templates/footer');
    }
}