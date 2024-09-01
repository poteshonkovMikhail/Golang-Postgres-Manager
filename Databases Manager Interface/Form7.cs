using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using static System.Windows.Forms.VisualStyles.VisualStyleElement;

namespace Databases_Manager_Interface
{
    public partial class Form7 : Form
    {
        public Form7()
        {
            InitializeComponent();
        }

        private async void button1_Click(object sender, EventArgs e)
        {
            string url = "http://" + textBox6.Text + ":8080/create-database";
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

        public static async Task<ConnectingDatabaseData> GetApiResponseAsync(string url, string data)
        {
            string jsonString = await FetchDataAsync(url, data);
            return JsonConvert.DeserializeObject<ConnectingDatabaseData>(jsonString);
        }

        private void button1_Click_1(object sender, EventArgs e)
        {
            Form7 form7 = new Form7();
            form7.Show();
        }
    }
}
