<?php namespace App\Controllers;

use App\Models\UsersModel;
use CodeIgniter\Controller;

class Users extends BaseController
{
    public function index()
    {
        if(isset($this->session->is_logged_in) && $this->session->is_logged_in){
            # They are already logged in and don't need to be here.
            return redirect()->route('/');
        }else{
           # They need an account to continue
           return redirect()->to(base_url('users/signup_required'));
        }
    }

    public function view($slug = null)
    {
        $model = new UsersModel();

        $data['users'] = $model->getUsers($slug);

        if (empty($data['users']))
        {
            throw new \CodeIgniter\Exceptions\PageNotFoundException('Cannot find the users item: '. $slug);
        }

        $data['title'] = $data['users']['title'];

        echo view('templates/header', $data);
        echo view('users/view', $data);
        echo view('templates/footer', $data);
    }

    public function login(){
        // if(isset($session->user_data['userId'])){
        if(isset($this->session->is_logged_in) && $this->session->is_logged_in){
            # They are already logged in and don't need to be here.
            return redirect()->route('/');
        }else{
            $data['title'] = 'Login';
            echo view('templates/header', $data);
            if($this->request->getMethod() === 'post'){
                # They are trying to login
                if($this->validate(['username' => 'required|min_length[1]',
                                    'password' => 'required'])){
                    # If their request was well formed.
                    $password = $this->request->getPost('password');
                    $username = $this->request->getPost('username');
                    $model = new UsersModel();
                    # Check if submitted credentials match a stored user
                    $id = $model->getUserId($username);
                    if($id !== NULL){
                        $user = $model->getUser($id);
                        if(password_verify($password, $user["password"])){
                            $user_data = ['username' => $username, 
                                      'id' => $model->getUserId($username),
                                      'is_logged_in' => TRUE];
                            $this->session->set($user_data);
                            return redirect()->route('/');
                        }else{
                            # It does not match a user so let them know
                            echo view('users/login', ['errors' => 'Invalid username or password.']);
                        }
                    }else{
                        # It does not match a user so let them know
                        echo view('users/login', ['errors' => 'Invalid username or password.']);
                    }
                }else{
                    # They submitted data that doesn't pass validation rules.
                    echo view('users/login', ['errors' => $this->validator->getErrors()]);
                }
            }else{
                # Display the login page
                echo view('users/login');
            }
            echo view('templates/footer');
        }
    }

    public function register(){
        if(isset($this->session->is_logged_in) && $this->session->is_logged_in){
            # They are already logged in and don't need to be here.
            return redirect()->route('/');
        }else{
            $data['title'] = 'Register';
            echo view('templates/header', $data);
            if($this->request->getMethod() === 'post'){
                # They are trying to login
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
                        # It does not match a user so they are good to signup.
                        $model->createUser($username, $password);
                        $user_data = ['username' => $username, 
                                      'id' => $model->getUserId($username),
                                      'is_logged_in' => TRUE];
                        $this->session->set($user_data);
                        return redirect()->route('/');
                    }
                }else{
                    # They submitted data that doesn't pass validation rules.
                    echo view('users/register', ['errors' => $this->validator->getErrors()]);
                }
            }else{
                # Display the login page
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