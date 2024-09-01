using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Databases_Manager_Interface
{
    using System;
    using System.Collections.Generic;
    using Newtonsoft.Json;

    public class TableInfo
    {
        [JsonProperty("table_name")]
        public string TableName { get; set; }

        [JsonProperty("columns")]
        public List<string> Columns { get; set; }

        [JsonProperty("rows")]
        public List<Dictionary<string, object>> Rows { get; set; }
    }

    public class CustomResponse
    {
        [JsonProperty("database_name")]
        public string DatabaseName { get; set; }

        [JsonProperty("tables")]
        public List<TableInfo> Tables { get; set; }
    }

    public class ConnectingDatabaseData
    {
        [JsonProperty("username_parameter")]
        public string UsernameParameter { get; set; }

        [JsonProperty("password_parameter")]
        public string PasswordParameter { get; set; }
        [JsonProperty("hostname_parameter")]
        public string HostnameParameter { get; set; }
        [JsonProperty("database_name_parameter")]
        public string DatabaseNameParameter { get; set; }

    }
    //public class SqlRequest
    //{
    //    [JsonProperty("username_parameter")]
    //    public string UsernameParameter { get; set; }

    //    [JsonProperty("password_parameter")]
    //    public string PasswordParameter { get; set; }
    //    [JsonProperty("hostname_parameter")]
    //    public string HostnameParameter { get; set; }
    //    [JsonProperty("database_name_parameter")]
    //    public string DatabaseNameParameter { get; set; }

    //    [JsonProperty("port_parameter")]
    //    public string PortParameter { get; set; }

    //    [JsonProperty("sqlCommand_parameter")]
    //    public string SqlCommand { get; set; }
    //}
    public class SqlResponse
    {
        [JsonProperty("results")]
        public List<Dictionary<string, object>> Results { get; set; } = new List<Dictionary<string, object>>();

    }

    public class DbInfo
    {
        [JsonProperty("database_id")]
        public int DatabaseId { get; set; }

        [JsonProperty("user_id")]
        public int UserId { get; set; }
        [JsonProperty("database_name")]
        public string DatabaseName { get; set; }
        [JsonProperty("database_password")]
        public string DatabasePassword { get; set; }
        [JsonProperty("database_hostname")]
        public string DatabaseHostname { get; set; }
        [JsonProperty("connected")]
        public bool Connected { get; set; }
        [JsonProperty("database_username")]
        public string DatabaseUsername { get; set; }

    }
}

