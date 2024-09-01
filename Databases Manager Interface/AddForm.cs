using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.Net.Http;
using System.Net.Http.Headers;
using Newtonsoft.Json;
using System.Xml.Linq;
using System.Threading;
using static System.Windows.Forms.VisualStyles.VisualStyleElement;


namespace Databases_Manager_Interface
{
    public partial class AddForm : Form
    {
        public AddForm()
        {
            InitializeComponent();
        }

        private void textBox4_TextChanged(object sender, EventArgs e)
        {

        }

        private async void button1_Click(object sender, EventArgs e)
        {
            string url = "http://" + textBox6.Text + ":8080/connectDB";
            // Создаем объект данных для отправки
            //var data = new
            //{
            //    // Можете изменить на свои нужные поля
            //    Username = textBox1.Text,
            //    Password = textBox3.Text,
            //    Hostname = textBox6.Text,
            //    DBName = textBox8.Text
            //};
            string h = textBox1.Text;
            string data = "{\"username_parameter\":\"" + h + "\"," +
                " \"password_parameter\":\"" + textBox3.Text + "\"," +
                " \"hostname_parameter\":\"" + textBox6.Text + "\"," +
                " \"database_name_parameter\":\"" + textBox8.Text + "\"}";
            try {
                ConnectingDatabaseData responseContent = await GetApiResponseAsync(url, data);

                if (responseContent.UsernameParameter == "This database already connected" &&
                    responseContent.PasswordParameter == "This database already connected" &&
                    responseContent.HostnameParameter == "This database already connected" &&
                    responseContent.DatabaseNameParameter == "This database already connected")
                {
                    textBox1.Text = "";
                    textBox3.Text = "";
                    textBox6.Text = "";
                    textBox8.Text = "";
                    textBox9.Text = "";
                    textBox10.Text = "";
                    textBox11.Text = "";
                    textBox12.Text = "";
                    MessageBox.Show("This database already connected");


                }
                if (responseContent.UsernameParameter == "V" &&
                    responseContent.PasswordParameter == "V" &&
                    responseContent.HostnameParameter == "V" &&
                    responseContent.DatabaseNameParameter == "V")
                {
                    Program.f1.dataGridView1.Rows.Add(textBox8.Text);
                    textBox1.Text = "";
                    textBox3.Text = "";
                    textBox6.Text = "";
                    textBox8.Text = "";
                }
                if (responseContent.UsernameParameter == "No such user" ||
                    responseContent.PasswordParameter == "Incorrect password" ||
                    responseContent.HostnameParameter == "No such hostname" ||
                    responseContent.DatabaseNameParameter == "No such database name")
                {
                    //string[] statuses = responseContent.Split(new char[] { '_' }, StringSplitOptions.RemoveEmptyEntries);
                    textBox9.Text = responseContent.UsernameParameter;
                    textBox10.Text = responseContent.PasswordParameter;
                    textBox11.Text = responseContent.HostnameParameter;
                    textBox12.Text = responseContent.DatabaseNameParameter;
                }



            }
            catch { 
            
                textBox11.Text = "Invalid hostname";
            }





        }

        private void textBox1_TextChanged(object sender, EventArgs e)
        {
            
        }

        static async Task<string> SendJsonRequestAsync(string url, object data)
        {
            using (HttpClient client = new HttpClient())
            {
                try
                {
                    // Сериализация объекта в JSON
                    string json = JsonConvert.SerializeObject(data);
                    var content = new StringContent(json, Encoding.UTF8, "application/json");

                    // Отправка POST-запроса с JSON-данными
                    HttpResponseMessage response = await client.PostAsync(url, content);

                    // Проверка успешности запроса
                    response.EnsureSuccessStatusCode();

                    // Чтение и возврат содержимого ответа
                    string responseContent = await response.Content.ReadAsStringAsync();
                    return responseContent;
                }
                catch (HttpRequestException e)
                {
                    Console.WriteLine($"Ошибка при отправке запроса: {e.Message}");
                    return null;
                }
            }
        }

        private static async Task<string> FetchDataAsync(string url, string data)
        {
            using (HttpClient client = new HttpClient())
            {
                var content = new StringContent(data, Encoding.UTF8, "application/json");
                HttpResponseMessage response = await client.PostAsync(url, content);
                response.EnsureSuccessStatusCode();
                return await response.Content.ReadAsStringAsync();
            }
        }

        private static async Task<ConnectingDatabaseData> GetApiResponseAsync(string url, string data)
        {
            string jsonString = await FetchDataAsync(url, data);
            return JsonConvert.DeserializeObject<ConnectingDatabaseData>(jsonString);
        }

        private void AddForm_Load(object sender, EventArgs e)
        {

        }
    }
}
